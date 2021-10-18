package auth

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/auth/hash"
	"github.com/sizata-siege/finance-management/auth/jwt"
	"github.com/sizata-siege/finance-management/user"
	"go.mongodb.org/mongo-driver/bson"
)

/* Requests */

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/* Auth middleware */

var Middleware = jwtware.New(jwtware.Config{
	ErrorHandler: authErrorHandler,
	SigningKey:   []byte(jwt.SECRET),
	TokenLookup:  "cookie:" + jwt.CookieName,
})

func TmpMiddleware(c *fiber.Ctx) error {
	/* Call basic middleware first */
	if err := Middleware(c); err != nil {
		return err
	} // maybe make a fun to parse user ... usr:=auth.user(c) ...

	return nil
	/* parse user claims */
	// usr := c.Locals("user").(*gojwt.Token)
}

/* handlers & controllers */

func Login(c *fiber.Ctx) error {
	usr := user.User{}
	var req LoginRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err, // todo: ok?
		})
	}

	/* check user & pass */
	err = mgm.Coll(&usr).First(bson.M{"email": req.Email}, &usr)

	/* check if user not found */
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if usr.Password != hash.GenerateHash(req.Password) {
		return fiber.ErrUnauthorized
	}

	t, err := jwt.GenerateToken(map[string]interface{}{
		"id":  usr.ID,
		"exp": jwt.DefaultSessionExpUnix(),
		// "name": usr.Name,
		// "admin": false,
	})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:    jwt.CookieName,
		Value:   t,
		Expires: jwt.DefaultSessionExp(),
	})
	return c.JSON(fiber.Map{
		"token": t,
		"id":    usr.ID,
		"name":  usr.Name,
		"email": usr.Email,
	})
}

func CreateNewUser(c *fiber.Ctx) error {
	usr := new(user.User)
	// usr := user.Create()
	err := c.BodyParser(usr)
	if err != nil {
		return fiber.ErrBadRequest
	}

	/* Generate hash */
	usr.Password = hash.GenerateHash(usr.Password)

	/* Save to db */
	err = mgm.Coll(usr).Create(usr)
	if err != nil {
		return err
	}

	// log the user in ...

	/* Remove password & respond */
	// usr.Password = ""
	t, err := jwt.GenerateToken(map[string]interface{}{
		"id":  usr.ID,
		"exp": jwt.DefaultSessionExpUnix(),
	})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":    usr.ID,
		"name":  usr.Name,
		"email": usr.Email,
		"token": t,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	usr := &user.User{}
	err := mgm.Coll(usr).FindByID(c.Params("id"), usr)
	if err != nil {
		return fiber.ErrNotFound
	}
	err = mgm.Coll(usr).Delete(usr)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func Logout(c *fiber.Ctx) error {
	return nil
}

func authErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   err,
		"message": "Unauthorized",
	})
}
