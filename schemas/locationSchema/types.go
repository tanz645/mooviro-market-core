package locationSchema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type geoLocation struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type CreateLocation struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Serial       uint32             `json:"serial" validate:"required"`
	Name         string             `json:"name" validate:"required"`
	Type         string             `json:"type" validate:"required"`
	ParentSerial uint32             `json:"parent_serial" bson:"parent_serial"`
	GeoLocation  geoLocation        `json:"geo_location" bson:"geo_location" validate:"required"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type SearchLocation struct {
	Limit        uint16 `form:"limit" validate:"required,max=500,min=1"`
	Page         uint16 `form:"page" validate:"required,min=1"`
	Type         string `form:"type"`
	Name         string `form:"name"`
	ParentSerial uint32 `form:"parent_serial"`
	// Serial       uint32 `form:"serial"`
	SortBy    string `form:"sort_by" validate:"required,oneof=name created_at"`
	SortOrder int8   `form:"sort_order" validate:"required,min=-1,max=1"`
}

type LocationGeneral struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Serial       uint32             `json:"serial"`
	Name         string             `json:"name"`
	Type         string             `json:"type"`
	ParentSerial uint32             `json:"parent_serial" bson:"parent_serial"`
	GeoLocation  geoLocation        `json:"geo_location" bson:"geo_location"`
}
