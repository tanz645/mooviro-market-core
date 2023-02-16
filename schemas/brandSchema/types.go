package brandSchema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateBrand struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name      string             `json:"name" validate:"required"`
	Logo      string             `json:"logo" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type BrandGeneral struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name string             `json:"name"`
	Logo string             `json:"logo"`
}
