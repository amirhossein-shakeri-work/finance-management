package db

import "github.com/kamva/mgm/v3"

type BaseModel mgm.Model

func (m BaseModel) Find(id string) *BaseModel {
	mgm.Coll(m)
	return nil
}
