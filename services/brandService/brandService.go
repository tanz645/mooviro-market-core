package brandService

import (
	"context"
	"matar/clients"
	"matar/schemas/brandSchema"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBrandById(ctx context.Context, id string) (*brandSchema.BrandGeneral, error) {
	var brand brandSchema.BrandGeneral
	var brandCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), brandSchema.BrandCollectionName)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "logo", Value: 1},
	}
	opts := options.FindOne().SetProjection(projection)
	err = brandCollection.FindOne(ctx, bson.D{{Key: "_id", Value: objId}}, opts).Decode(&brand)
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func GetBrands(ctx context.Context) ([]brandSchema.BrandGeneral, error) {
	var brands []brandSchema.BrandGeneral
	var brandCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), brandSchema.BrandCollectionName)

	sort := bson.D{{Key: "name", Value: 1}}
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "logo", Value: 1},
	}
	opts := options.Find().SetProjection(projection).SetSort(sort)
	results, err := brandCollection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	for results.Next(ctx) {
		var brand brandSchema.BrandGeneral
		if err = results.Decode(&brand); err != nil {
			return nil, err
		}

		brands = append(brands, brand)
	}
	return brands, nil
}
