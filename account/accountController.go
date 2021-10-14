package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func IndexAccounts (c *fiber.Ctx) error {
	var accounts []Account
	err := mgm.Coll(&Account{}).SimpleFind(&accounts, bson.M{})
	if err != nil { return err }
	return c.JSON(accounts)
}

func ShowAccount (c *fiber.Ctx) error {
	acc := &Account{}
	err := mgm.Coll(acc).FindByID(c.Params("id"), acc)
	if err != nil { return err }
	return c.JSON(acc)
	// return fiber.ErrNotFound
}

func StoreAccount (c *fiber.Ctx) error {
	acc := new(Account)
	err := c.BodyParser(acc)
	if err != nil { return err }
	err = mgm.Coll(acc).Create(acc)
	if err != nil { return err }
	return c.Status(fiber.StatusCreated).JSON(acc)
}

func UpdateAccount (c *fiber.Ctx) error {
	return c.SendString("Update Account")
}

func DeleteAccount (c *fiber.Ctx) error {
	// fixme: returns 500 if nothing found, send 404 instead
	acc := &Account{}
	err := mgm.Coll(acc).FindByID(c.Params("id"), acc)
	if err != nil { return err }
	err = mgm.Coll(acc).Delete(acc)
	if err != nil { return err }
	c.Status(fiber.StatusNoContent)
	return nil
	// return fiber.ErrNotFound
}
