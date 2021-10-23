package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/account"
	"github.com/sizata-siege/finance-management/auth/jwt"
)

func IndexAccounts (c *fiber.Ctx) error {
	return c.JSON(jwt.New(c).User.Accounts)
	// var accounts []account.Account
	// err := mgm.Coll(&account.Account{}).SimpleFind(&accounts, bson.M{})
	// if err != nil { return err }
	// return c.JSON(accounts)
}

func ShowAccount (c *fiber.Ctx) error {
	acc := &account.Account{}
	err := mgm.Coll(acc).FindByID(c.Params("id"), acc)
	if err != nil { return err }
	return c.JSON(acc)
	// return fiber.ErrNotFound
}

func StoreAccount (c *fiber.Ctx) error {
	acc := new(account.Account)
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
	acc := &account.Account{}
	err := mgm.Coll(acc).FindByID(c.Params("id"), acc)
	if err != nil { return err }
	err = mgm.Coll(acc).Delete(acc)
	if err != nil { return err }
	c.Status(fiber.StatusNoContent)
	return nil
	// return fiber.ErrNotFound
}
