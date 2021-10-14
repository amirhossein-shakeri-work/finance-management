package session

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sizata-siege/finance-management/user"
)

const SECRET = "Fuck USA"

/* Requests */
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/* Auth middleware */

var Auth = jwtware.New(jwtware.Config{
	// ErrorHandler: authErrorHandler,
	SigningKey: []byte(SECRET),
})

func Signup(c *fiber.Ctx) error {
	return nil
}

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
	match := false
	if !match {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	/* Create token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Set claims */
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "name"
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	/* Generate encoded token and send it as response */
	t, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": t,
		"user": usr,
	})
}

func Logout(c *fiber.Ctx) error {
	return nil
}

func authErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}
