package automobileAdSchema

import (
	"context"
	"fmt"
	"matar/clients"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var AutomobileAdCollectionName = "automobile_ad"

func CreateAutomobileAdIndexes(ctx context.Context, client *mongo.Client) {
	col := clients.GetMongoCollection(client, AutomobileAdCollectionName)
	models := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "address.id", Value: bsonx.Int32(1)},
				{Key: "body_type", Value: bsonx.Int32(1)},
				{Key: "brand.id", Value: bsonx.Int32(1)},
				{Key: "transmission", Value: bsonx.Int32(1)},
				{Key: "fuel_type", Value: bsonx.Int32(1)},
				{Key: "wheel_drive", Value: bsonx.Int32(1)},
				{Key: "price.total_amount", Value: bsonx.Int32(1)},
				{Key: "milage.amount", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "body_type", Value: bsonx.Int32(1)},
				{Key: "address.id", Value: bsonx.Int32(1)},
				{Key: "brand.id", Value: bsonx.Int32(1)},
				{Key: "transmission", Value: bsonx.Int32(1)},
				{Key: "fuel_type", Value: bsonx.Int32(1)},
				{Key: "wheel_drive", Value: bsonx.Int32(1)},
				{Key: "price.total_amount", Value: bsonx.Int32(1)},
				{Key: "milage.amount", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "brand.id", Value: bsonx.Int32(1)},
				{Key: "address.id", Value: bsonx.Int32(1)},
				{Key: "body_type", Value: bsonx.Int32(1)},
				{Key: "transmission", Value: bsonx.Int32(1)},
				{Key: "fuel_type", Value: bsonx.Int32(1)},
				{Key: "wheel_drive", Value: bsonx.Int32(1)},
				{Key: "price.total_amount", Value: bsonx.Int32(1)},
				{Key: "milage.amount", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "transmission", Value: bsonx.Int32(1)},
				{Key: "address.id", Value: bsonx.Int32(1)},
				{Key: "body_type", Value: bsonx.Int32(1)},
				{Key: "brand.id", Value: bsonx.Int32(1)},
				{Key: "fuel_type", Value: bsonx.Int32(1)},
				{Key: "wheel_drive", Value: bsonx.Int32(1)},
				{Key: "price.total_amount", Value: bsonx.Int32(1)},
				{Key: "milage.amount", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "user_id", Value: bsonx.Int32(1)},
				{Key: "created_at", Value: bsonx.Int32(1)},
			},
		},
	}
	_, err := col.Indexes().CreateMany(ctx, models)
	if err != nil {
		fmt.Println(err)
	}
}
