package programmesneufs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
how to convert ObjectID:
- declare type = primitive.ObjectID from bson/primitive
- use tag declaration bson:"_id"
- json tags only for mapping presentation
- if null field: put *string
*/
type Geojson struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type Programmesneufs struct {
	ID                  primitive.ObjectID       `bson:"_id,omitempty" json:"_id,omitempty"`
	Name                *string                  `bson:"name,omitempty" json:"name,omitempty"`
	Description         *string                  `json:"description,omitempty"`
	Price               interface{}              `bson:"price,omitempty" json:"price,omitempty"`
	CreationDate        string                   `bson:"creationDate,omitempty" json:"created_at,omitempty"`
	Coordinates         map[string]interface{}   `bson:"coordinates,omitempty" json:"coordinates,omitempty"`
	ProfessionalLogoUrl map[string]interface{}   `bson:"professionalLogoUrl,omitempty" json:"professionalLogoUrl,omitempty"`
	ProfessionalName    interface{}              `bson:"professionalName,omitempty" json:"professionalName,omitempty"`
	Thumbnail           []map[string]interface{} `bson:"thumbnail,omitempty" json:"thumbnail,omitempty"`
	Geojson             Geojson                  `bson:"geojson,omitempty" json:"geojson,omitempty"`
}

/**
  id: string;
  description?: string;
  logo?: string;
  name?: string;
  onClick?: (event: MouseEvent) => void;
  title?: string;
  price?: number;
  // eslint-disable-next-line
  thumbnails?: any;
*/
