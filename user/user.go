package user

import (
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/account"
	"github.com/sizata-siege/finance-management/auth/hash"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/tag"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string            `json:"name" bson:"name"`
	Email            string            `json:"email" bson:"email"`
	Password         string            `json:"password" bson:"password"`
	Tags             tag.Set           `json:"tags" bson:"tags"`
}

type Attr struct {
	name     string
	email    string
	password string
}

func New(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Tags: tag.Set{},
	}
}

func Create(a Attr) (*User, error) { // is it ok to return pointer?
	usr := New(a.name, a.email, a.password)

	/* hash password */
	usr.Password = hash.GenerateHash(usr.Password)

	/* save to db */
	return usr, mgm.Coll(usr).Create(usr)
}

func (user *User) Delete() (*User, error) {
	return user, mgm.Coll(user).Delete(user)
}

/* =-=-=-=-=-=-= DB Helpers =-=-=-=-=-=-= */

func Find(id string) *User {
	u := &User{}
	if err := mgm.Coll(u).FindByID(id, u); err != nil {
		return nil
	}
	return u
}

func FindByEmail(email string) *User {
	u := &User{}
	if err := mgm.Coll(u).First(bson.M{"email": email}, u); err != nil { // get all of them?
		return nil
	}
	return u
}

func (user *User) Save () error { // could return *User
	if err := mgm.Coll(user).Update(user); err != nil {
		return err
	}
	return nil
}

/* =-=-=-=-=-=-= Relations =-=-=-=-=-=-= */

func (user *User) Accounts () []*account.Account {
	var accounts []*account.Account
	if err := mgm.Coll(&account.Account{}).SimpleFind(accounts, bson.M{"user_id": user.ID}); err != nil {
		log.Println(err)
		return nil
	}
	return accounts
}

// func (user *User) AddAccount (acc *account.Account) error {
// 	// Add the passed account
// 	return nil
// }

// func (user *User) CreateAccount (name string, balance float64) (*account.Account, error) {
// 	/* Create & Save account in user accounts */
// 	acc := account.NewWithID(name, balance)
// 	if err := acc.Creating(); err != nil {
// 		return nil, err
// 	}
// 	if user.Accounts == nil {
// 		user.Accounts = []account.Account{}
// 	}
// 	user.Accounts = append(user.Accounts, *acc)
// 	if err := user.Save(); err != nil {
// 		return nil, err
// 	}
// 	return acc, nil
// }

// func (user *User) RemoveAccount (id string) error {
// 	/* find account */
// 	for i, acc := range user.Accounts {
// 		/* Compare ids */
// 		if acc.ID.Hex() == id {
// 			user.Accounts = append(user.Accounts[:i], user.Accounts[i + 1:]...)
// 			return user.Save()
// 		}
// 	}
// 	return errors.New("account not found " + id)
// }

// func (user *User) GetAccount (id string) *account.Account {
// 	/* find account */
// 	for _, acc := range user.Accounts {
// 		if acc.ID.Hex() == id {
// 			return &acc
// 		}
// 	}
// 	/* Not Found */
// 	return nil
//
// 	/* by reference?! */
// 	// for i, _ := range user.Accounts {
// 	// 	if user.Accounts[i].ID.Hex() == id {
// 	// 		return &user.Accounts[i]
// 	// 	}
// 	// }
// }
