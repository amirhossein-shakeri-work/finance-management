package account

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/tag"
)

type Account struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	Balance          float64            `json:"balance" bson:"balance"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	Tags             tag.Set            `json:"tags" bson:"tags"`
} // default model adds _id, timestamps. inline flattens the struct

type Attr struct {
	name    string
	balance float64
	userId  primitive.ObjectID
}

func New(name string, balance float64, userId primitive.ObjectID) *Account {
	return &Account{
		Name:    name,
		Balance: balance,
		UserID: userId,
	}
}

func Create(a Attr) (*Account, error) {
	acc := New(a.name, a.balance, a.userId)
	return acc, mgm.Coll(acc).Create(acc)
}

func (acc *Account) Delete () (*Account, error) {
	return acc, mgm.Coll(acc).Delete(acc)
}
