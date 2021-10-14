package transaction

import "github.com/kamva/mgm/v3"

type Transaction struct {
	mgm.DefaultModel `bson:",inline"`
	Type             string  `json:"type" bson:"type"`
	Amount           float64 `json:"amount" bson:"amount"`
	AccountID mgm.IDField
}
