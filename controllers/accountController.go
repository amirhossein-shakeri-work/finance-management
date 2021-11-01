package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sizata-siege/finance-management/auth/jwt"
)

/* Requests */

type CreateAccountRequest struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func IndexAccounts(c *fiber.Ctx) error {
	return c.JSON(jwt.New(c).User.Accounts)
	// var accounts []account.Account
	// err := mgm.Coll(&account.Account{}).SimpleFind(&accounts, bson.M{})
	// if err != nil { return err }
	// return c.JSON(accounts)
}

func ShowAccount(c *fiber.Ctx) error {
	/* Access User */
	user := jwt.New(c).User

	/* Retrieve Account */
	acc := user.GetAccount(c.Params("id"))
	if acc == nil {
		return fiber.ErrNotFound
	}
	return c.JSON(acc)
}

func CreateAccount(c *fiber.Ctx) error {
	req := CreateAccountRequest{} // name, balance
	/* Parse Request Body */
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	/* Access User */
	user := jwt.New(c).User
	acc, err := user.CreateAccount(req.Name, req.Balance)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(acc)
}

func UpdateAccount(c *fiber.Ctx) error {
	return c.SendString("Update Account Coming Soon")
}

func DeleteAccount(c *fiber.Ctx) error {
	/* Access User */
	user := jwt.New(c).User
	if err := user.RemoveAccount(c.Params("id")); err != nil {
		/* Not Found */
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
