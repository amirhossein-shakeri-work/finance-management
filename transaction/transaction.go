package transaction

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/tag"
)

type Transaction struct {
	mgm.DefaultModel `bson:",inline"`
	Type             string             `json:"type" bson:"type"`
	Amount           float64            `json:"amount" bson:"amount"`
	Description      string             `json:"description" bson:"description"`
	AccountID        primitive.ObjectID `json:"account_id" bson:"account_id"`
	Tags             tag.Set            `json:"tags" bson:"tags"`
}

type Attr struct {
	t      string
	amount float64
	accId  primitive.ObjectID
}

func New(t string, amount float64, accId primitive.ObjectID) *Transaction {
	return &Transaction{
		Type:      t,
		Amount:    amount,
		AccountID: accId,
	}
}

func Create(a Attr) (*Transaction, error) {
	tr := New(a.t, a.amount, a.accId)
	return tr, mgm.Coll(tr).Create(tr)
}

func (tr *Transaction) Delete() (*Transaction, error) {
	return tr, mgm.Coll(tr).Delete(tr)
}
