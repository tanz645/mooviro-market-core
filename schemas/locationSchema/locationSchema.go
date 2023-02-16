package locationSchema

import (
	"context"
	"fmt"
	"matar/clients"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var LocationCollectionName = "locations"

func CreateLocationIndexes(ctx context.Context, client *mongo.Client) {
	col := clients.GetMongoCollection(client, LocationCollectionName)
	models := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "name", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "type", Value: bsonx.Int32(1)},
				{Key: "parent_serial", Value: bsonx.Int32(1)},
			},
		},
	}
	_, err := col.Indexes().CreateMany(ctx, models)
	if err != nil {
		fmt.Println(err)
	}
}
