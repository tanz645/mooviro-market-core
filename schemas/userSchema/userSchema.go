package userSchema

import (
	"context"
	"matar/clients"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollectionName = "users"

func CreateUserIndexes(ctx context.Context, client *mongo.Client) {
	col := clients.GetMongoCollection(client, UserCollectionName)
	col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"phone": 1},
	})
}
