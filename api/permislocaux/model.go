package permislocaux

import "github.com/mongodb/mongo-go-driver/bson/primitive"


/**
how to convert ObjectID:
- declare type = primitive.ObjectID from bson/primitive
- use tag declaration bson:"_id"
- json tags only for mapping presentation
*/
type Permislocaux struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Siret                  string             `json:"siret"`
	Codepostaledudemandeur string             `string:"codepostal"`
}
