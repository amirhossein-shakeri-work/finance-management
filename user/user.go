package user

import (
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/auth/hash"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
}

func New (name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func Create (name, email, password string) (*User, error) { // is it ok to return pointer?
	usr := New(name, email, password)

	/* hash password */
	usr.Password = hash.GenerateHash(usr.Password)

	/* save to db */
	err := mgm.Coll(usr).Create(usr)
	return usr, err
}

func (user *User) Update (attr map[string]interface{}) (*User, error) {
	// for key, value := range attr {
	// 	user
	// }
	return user, nil
}

// func deleteAccount(c *fiber.Ctx) error {
// 	return nil
// }
