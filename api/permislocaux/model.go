package permislocaux

import "go.mongodb.org/mongo-driver/bson/primitive"


/**
how to convert ObjectID:
- declare type = primitive.ObjectID from bson/primitive
- use tag declaration bson:"_id"
- json tags only for mapping presentation
*/
type Permislocaux struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Siret                  string             `json:"siret"`
	Codepostaledudemandeur string             `json:"codepostal"`
}
