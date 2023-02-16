package brandSchema

import (
	"context"
	"fmt"
	"matar/clients"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var BrandCollectionName = "brands"

func CreateBrandIndexes(ctx context.Context, client *mongo.Client) {
	col := clients.GetMongoCollection(client, BrandCollectionName)
	models := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "name", Value: bsonx.Int32(1)},
			},
		},
	}
	_, err := col.Indexes().CreateMany(ctx, models)
	if err != nil {
		fmt.Println(err)
	}
}
