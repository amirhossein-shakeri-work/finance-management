package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/transaction"
)

/* Requests */

type CreateTransactionRequest struct {
	Source      string  `json:"source"` // primitive.ObjectID
	Amount      float64 `json:"amount"`
	Destination string  `json:"destination"` // primitive.ObjectID
	Description string  `json:"description"`
}

func IndexTransactions(c *fiber.Ctx) error {
	/* Get User */
	/* Get Trans based on filters passed (src, dest, amount, id, all_of_user) */
	return nil
}

func CreateTransaction(c *fiber.Ctx) error {
	req := CreateTransactionRequest{}

	/* Parse Request Body */
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	tr := transaction.New(req.Source, req.Amount, req.Destination, req.Description)

	/* Validate Transaction */
	if isValid, err := tr.Validate(); !isValid {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	// isValid, err := tr.Validate()
	// fmt.Printf("%+v\n%+v\n%+v\n%+v\n", tr, isValid, err, req)

	/* Create Transaction */
	if err := mgm.Coll(tr).Create(tr); err != nil {
		return err
	}

	/* Decrease source & Increase Destination by amount */
	// how to find the account with given id in src or dest?
	// we can find the user contains the acc with target accID
	// and then do the job
	// , or
	// we can get accounts in a separate collection! what to do now???

	return nil
}
