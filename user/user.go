package user

import (
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/auth/hash"
	"go.mongodb.org/mongo-driver/tag"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string  `json:"name"`
	Email            string  `json:"email"`
	Password         string  `json:"password"`
	Tags             tag.Set `json:"tags" bson:"tags"`
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
