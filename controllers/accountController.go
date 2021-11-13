package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/account"
	"github.com/sizata-siege/finance-management/auth/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

/* Requests */

type CreateAccountRequest struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func IndexAccounts(c *fiber.Ctx) error {
	user := jwt.New(c).User
	var accounts []account.Account
	err := mgm.Coll(&account.Account{}).SimpleFind(&accounts, bson.M{"user_id": user.ID})
	if err != nil { return err }
	return c.JSON(accounts)
	// return c.JSON(jwt.New(c).User.Accounts)
}

func ShowAccount(c *fiber.Ctx) error {
	acc := &account.Account{}
	if err := mgm.Coll(acc).FindByID(c.Params("id"), acc); err != nil {
		return err
	}
	if jwt.New(c).User.ID != acc.ID {
		return fiber.ErrForbidden // Those bitches only have access to their own accounts!
	}
	return c.JSON(acc)
}

func CreateAccount(c *fiber.Ctx) error {
	req := CreateAccountRequest{} // name, balance
	/* Parse Request Body */
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	/* Access User & Create Account */
	user := jwt.New(c).User
	acc, err := account.Create(user.ID, req.Name, req.Balance)
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
	acc := account.Find(c.Params("id"))
	if user.ID != acc.UserID {
		return fiber.ErrForbidden // U can't delete others' accounts, bitch!
	}
	if err := acc.Delete(); err != nil {
		/* Not Found */
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
