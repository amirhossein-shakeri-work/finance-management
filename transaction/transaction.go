package transaction

import (
	"errors"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/tag"
)

type Transaction struct {
	mgm.DefaultModel `bson:",inline"`
	Source           string  `json:"source" bson:"source"` // or primitive.ObjectID
	Amount           float64 `json:"amount" bson:"amount"`
	Destination      string  `json:"destination" bson:"destination"` // or primitive.ObjectID
	Description      string  `json:"description" bson:"description"`
	Tags             tag.Set `json:"tags" bson:"tags"`
	// Type             string             `json:"type" bson:"type"` // no idea
}

type Attr struct {
	source      string // primitive.ObjectID
	amount      float64
	destination string // primitive.ObjectID
	description string
}

func New(source string, amount float64, destination, desc string) *Transaction {
	return &Transaction{
		Source:      source,
		Amount:      amount,
		Destination: destination,
		Description: desc,
	}
}

/* =-=-=-=-=-=-= DB Helpers =-=-=-=-=-=-= */

func Create(a Attr) (*Transaction, error) {
	tr := New(a.source, a.amount, a.destination, a.description)
	return tr, mgm.Coll(tr).Create(tr)
}

func (tr *Transaction) Delete() (*Transaction, error) {
	return tr, mgm.Coll(tr).Delete(tr)
}

func Find(id string) *Transaction {
	t := &Transaction{}
	if err := mgm.Coll(t).FindByID(id, t); err != nil {
		return nil
	}
	return t
}

func (tr *Transaction) Save() error {
	if err := mgm.Coll(tr).Update(tr); err != nil {
		return err
	}
	return nil
}

func (tr *Transaction) Validate() (bool, error) {
	/* At least one of source or destination must be valid */
	if !primitive.IsValidObjectID(tr.Source) && !primitive.IsValidObjectID(tr.Destination) {
		return false, errors.New("invalid source or destination")
	}

	/* Source & Destination can't be the same */
	if tr.Source == tr.Destination {
		return false, errors.New("source and destination can't be the same")
	}

	return true, nil
}
