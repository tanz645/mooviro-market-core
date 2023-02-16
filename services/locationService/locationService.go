package locationService

import (
	"context"
	"matar/clients"
	"matar/schemas/locationSchema"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetLocationById(ctx context.Context, id string) (*locationSchema.LocationGeneral, error) {
	var location locationSchema.LocationGeneral
	var locationCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), locationSchema.LocationCollectionName)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "serial", Value: 1},
		{Key: "type", Value: 1},
		{Key: "parent_serial", Value: 1},
		{Key: "geo_location", Value: 1},
	}
	opts := options.FindOne().SetProjection(projection)
	err = locationCollection.FindOne(ctx, bson.D{{Key: "_id", Value: objId}}, opts).Decode(&location)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func GetLocationsByParentSerial(ctx context.Context, serial uint32) ([]locationSchema.LocationGeneral, error) {
	var locations []locationSchema.LocationGeneral
	var locationCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), locationSchema.LocationCollectionName)
	sort := bson.D{{Key: "name", Value: 1}}
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "serial", Value: 1},
		{Key: "type", Value: 1},
		{Key: "parent_serial", Value: 1},
		{Key: "geo_location", Value: 1},
	}
	opts := options.Find().SetProjection(projection).SetSort(sort)
	results, err := locationCollection.Find(ctx, bson.D{
		{Key: "parent_serial", Value: serial},
	}, opts)
	if err != nil {
		return nil, err
	}
	for results.Next(ctx) {
		var location locationSchema.LocationGeneral
		if err = results.Decode(&location); err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}
	return locations, nil
}
