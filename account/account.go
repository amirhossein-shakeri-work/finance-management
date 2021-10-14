package account

import (
	"github.com/kamva/mgm/v3"
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

