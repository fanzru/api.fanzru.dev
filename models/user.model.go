package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	// nama variable, jenisnya, tag
	Name      string `json:"name" bson:"name"`
	Author    string `json:"author" bson:"author"`
}

