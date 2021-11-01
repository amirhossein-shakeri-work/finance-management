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
	Tags             tag.Set            `json:"tags" bson:"tags"`
} // default model adds _id, timestamps. inline flattens the struct

type Attr struct {
	name    string
	balance float64
}

func New(name string, balance float64) *Account {
	return &Account{
		Name:    name,
		Balance: balance,
	}
}

func NewWithID (name string, balance float64) *Account {
	acc := New(name, balance)
	acc.ID = primitive.NewObjectID()
	return acc
}
