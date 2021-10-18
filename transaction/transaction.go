package transaction

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/tag"
)

type Transaction struct {
	mgm.DefaultModel `bson:",inline"`
	Source           primitive.ObjectID `json:"source" bson:"source"` // or string
	Amount           float64            `json:"amount" bson:"amount"`
	Destination      primitive.ObjectID `json:"destination" bson:"destination"` // or string
	Description      string             `json:"description" bson:"description"`
	Tags             tag.Set            `json:"tags" bson:"tags"`
	// Type             string             `json:"type" bson:"type"` // no idea
}

type Attr struct {
	source      primitive.ObjectID
	amount      float64
	destination primitive.ObjectID
	description string
}

func New(source primitive.ObjectID, amount float64, destination primitive.ObjectID, desc string) *Transaction {
	return &Transaction{
		Source:      source,
		Amount:      amount,
		Destination: destination,
		Description: desc,
	}
}

func Create(a Attr) (*Transaction, error) {
	tr := New(a.source, a.amount, a.destination, a.description)
	return tr, mgm.Coll(tr).Create(tr)
}

func (tr *Transaction) Delete() (*Transaction, error) {
	return tr, mgm.Coll(tr).Delete(tr)
}
