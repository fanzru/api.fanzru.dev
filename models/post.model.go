package models
import (
	"github.com/kamva/mgm/v3"
)

type Post struct {
	mgm.DefaultModel `bson:",inline"`
	// nama variable, jenisnya, tag
	Name    	string `json:"name" bson:"name"`
	Message		string `json:"message" bson:"message"`
}
