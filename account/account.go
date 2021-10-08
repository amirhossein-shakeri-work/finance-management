package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Account struct {
	mgm.DefaultModel `bson:",inline"` // adds _id, timestamps. inline flattens the struct
	Name             string           `json:"name" bson:"name"`
	Balance          float64          `json:"balance" bson:"balance"`
}

func NewAccount(name string, balance float64) *Account {
	return &Account{
		Name:    name,
		Balance: balance,
	}
}

// var accounts = []Account{
// 	{Name: "Investing Account", Balance: 1000000000.24},
// 	{Name: "Temp Account", Balance: 1000000.56},
// }

func Index(c *fiber.Ctx) error {
	accounts := []Account{}
	err := mgm.Coll(&Account{}).SimpleFind(&accounts, bson.M{})
	if err != nil { return err }
	return c.JSON(accounts)
}

func Show(c *fiber.Ctx) error {
	acc := &Account{}
	err := mgm.Coll(acc).FindByID(c.Params("id"), acc)
	if err != nil { return err }
	return c.JSON(acc)
	// return fiber.ErrNotFound
}

func Store(c *fiber.Ctx) error {
	acc := new(Account)
	err := c.BodyParser(acc)
	if err != nil { return err }
	err = mgm.Coll(acc).Create(acc)
	if err != nil { return err }
	// accounts = append(accounts, *account)
	return c.Status(fiber.StatusCreated).JSON(acc)
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Update Account")
}

func Delete(c *fiber.Ctx) error {
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
