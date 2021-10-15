package user

import (
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/account"
	"github.com/sizata-siege/finance-management/auth/hash"
	"go.mongodb.org/mongo-driver/tag"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string            `json:"name" bson:"name"`
	Email            string            `json:"email" bson:"email"`
	Password         string            `json:"password" bson:"password"`
	Accounts         []account.Account `json:"accounts" bson:"accounts"` // https://github.com/Kamva/mgm/discussions/39
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
	}
}

func Create(a Attr) (*User, error) { // is it ok to return pointer?
	usr := New(a.name, a.email, a.password)

	/* hash password */
	usr.Password = hash.GenerateHash(usr.Password)

	/* save to db */
	return usr, mgm.Coll(usr).Create(usr)
}

// func (user *User) Update(a map[string]interface{}) (*User, error) { // map[string]interface{}
// 	// for key, value := range a {
// 	// 	switch key {
// 	// 	case "name": user.Name = string(value)
// 	// 	}
// 	// }
// 	// mgm.Coll(user).Update()
// 	return user, nil
// }

func (user *User) Delete() (*User, error) {
	return user, mgm.Coll(user).Delete(user)
}

/* Override Hooks to add id & date times for nested documents */
// https://github.com/Kamva/mgm/discussions/39

func (user *User) Creating() error {
	// now := time.Now().UTC()

	if err := user.DefaultModel.Creating(); err != nil {
		return err
	}

	for _, acc := range user.Accounts {
		if err := acc.Creating(); err != nil {
			return err
		}
	}
	return nil
}

func (user *User) Saving() error {
	if err := user.DefaultModel.Saving(); err != nil {
		return err
	}

	for _, acc := range user.Accounts {
		if err := acc.Saving(); err != nil {
			return err
		}
	}
	return nil
}
