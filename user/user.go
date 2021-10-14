package user

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func CreateNewUser(c *fiber.Ctx) error {
	usr := new(User)
	err := c.BodyParser(usr)
	if err != nil {
		return fiber.ErrBadRequest
	}

	/* Generate hash */
	h := sha512.Sum384([]byte(usr.Password))
	usr.Password = hex.EncodeToString(h[:])
	// h := sha512.New384()
	// h.Write([]byte(usr.Password))
	// usr.Password = string(h.Sum(nil))
	// usr.Password = string(sha512.Sum384([]byte(usr.Password))) // doesn't work?

	/* Save to db */
	err = mgm.Coll(usr).Create(usr)
	if err != nil {
		return err
	}

	// log the user in ...
	return c.Status(fiber.StatusCreated).JSON(usr)
}

func DeleteUser(c *fiber.Ctx) error {
	usr := &User{}
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

// func deleteAccount(c *fiber.Ctx) error {
// 	return nil
// }
