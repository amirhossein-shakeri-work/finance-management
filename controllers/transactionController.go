package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/account"
	"github.com/sizata-siege/finance-management/auth/jwt"
	"github.com/sizata-siege/finance-management/transaction"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Requests */

type IndexTransactionsRequest struct {
	AccID       primitive.ObjectID `json:"acc_id"`
	Source      string             `json:"source"`
	Destination string             `json:"destination"`
}

type CreateTransactionRequest struct {
	Source      string  `json:"source"` // primitive.ObjectID
	Amount      float64 `json:"amount"`
	Destination string  `json:"destination"` // primitive.ObjectID
	Description string  `json:"description"`
}

func IndexTransactions(c *fiber.Ctx) error {
	/* Get User */
	user := jwt.New(c).User

	/* if accId is passed in params like this /:accID/transactions */
	if accId := c.Params("id"); accId != "" {
		fmt.Println("accID passed in params: ", accId)
		acc := account.Find(accId)
		if acc == nil {
			return fiber.ErrNotFound
		}
		if acc.UserID != user.ID {
			return fiber.ErrForbidden // they can't fetch others' accounts!
		}

		/* find all transactions related to the account (source, destination) */
		if trs, err := transaction.RelatedToAccount(acc); err != nil {
			return err
		} else {
			fmt.Println("Done with result: ", trs)
			return c.JSON(trs)
		}
	}

	/* If no accId is passed in params, parse request body */
	req := &IndexTransactionsRequest{}
	if err := c.BodyParser(req); err != nil {
		return err
	}

	/* Get Trans based on filters passed (src, dest, amount, id, all_of_user) */
	/* Or get all transactions of all user accounts using cool db aggregations! */
	/* they can also filter using balance e.g. balance: {operator: '>=', value: 102} or {57: '<', 110.4: '>='} */
	/* also filter by date */
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

	/* Create Transaction */
	if err := mgm.Coll(tr).Create(tr); err != nil {
		return err
	}

	/* Decrease source & Increase Destination by amount */
	// we could use db transactions to undo the changes in accounts if somewhere not ok
	// affectsSrc := false
	// affectsDest := false
	if tr.HasValidSource() {
		/* Check if account balance is not negative */
		src := tr.SourceAcc()
		sub := src.Balance - tr.Amount
		if sub < 0 {
			/* Check if source account can have negative values */
			if !src.CanHaveNegativeBalance() {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "insufficient funds"})
			}
		}
		/* Apply source transaction effect */
		// affectsSrc = true
	}

	if tr.HasValidDestination() {
		/* Check for any errors to throw */
		/* Apply source transaction effect */
		// affectsDest = true
	}

	/* All errors are now checked & here nothing is wrong no need for transactions :) */
	if err := tr.Apply(); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(tr)
}

func UpdateTransaction(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func DeleteTransaction(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func UndoTransaction(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}
