package account

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/tag"
)

type Account struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name             string             `json:"name" bson:"name"`
	Balance          float64            `json:"balance" bson:"balance"`
	Tags             tag.Set            `json:"tags" bson:"tags"`
} // default model adds _id, timestamps. inline flattens the struct

// type Person struct {
// 	Name   string  `json:"name" bson:"name"`
// 	Age    int     `json:"age" bson:"age"`
// 	Weight float64 `json:"weight" bson:"weight"`
// }

type Attr struct {
	name    string
	balance float64
}

func New(uId primitive.ObjectID, name string, balance float64) *Account {
	return &Account{
		UserID:  uId,
		Name:    name,
		Balance: balance,
	}
}

func Create(uId primitive.ObjectID, name string, balance float64) (*Account, error) {
	acc := New(uId, name, balance)
	return acc, mgm.Coll(acc).Create(acc)
}

func Find(id string) *Account {
	acc := &Account{}
	if err := mgm.Coll(acc).FindByID(id, acc); err != nil {
		return nil
	}
	return acc
}

func Delete(id string) error {
	return Find(id).Delete()
}

func (acc *Account) Save() error {
	if err := mgm.Coll(acc).Update(acc); err != nil {
		return err
	}
	return nil
}

func (acc *Account) Delete() error {
	return mgm.Coll(acc).Delete(acc)
}

func (acc *Account) CanHaveNegativeBalance() bool {
	/* Check tags maybe? */
	return false
}

func (acc *Account) IncreaseBalance(n float64) error {
	acc.Balance += n
	return acc.Save()
}

func (acc *Account) DecreaseBalance(n float64) error {
	return acc.IncreaseBalance(-n)
}

// func (acc *Account) Validate() (bool, error) {
// 	/* Check if UserId is valid & existing */
// }

// func NewWithID(name string, balance float64) *Account {
// 	acc := New(name, balance)
// 	acc.ID = primitive.NewObjectID()
// 	return acc
// }
