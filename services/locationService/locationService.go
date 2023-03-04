package locationService

import (
	"context"
	"matar/clients"
	"matar/common/responses"
	"matar/schemas/locationSchema"
	"strings"

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

func GetLocationByParentSerial(ctx context.Context, serial uint32) (*locationSchema.LocationGeneral, error) {
	var location locationSchema.LocationGeneral
	var locationCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), locationSchema.LocationCollectionName)
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "serial", Value: 1},
		{Key: "type", Value: 1},
		{Key: "parent_serial", Value: 1},
		{Key: "geo_location", Value: 1},
	}
	opts := options.FindOne().SetProjection(projection)
	err := locationCollection.FindOne(ctx, bson.D{{Key: "parent_serial", Value: serial}}, opts).Decode(&location)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func GetLocationBySerial(ctx context.Context, serial uint32) (*locationSchema.LocationGeneral, error) {
	var location locationSchema.LocationGeneral
	var locationCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), locationSchema.LocationCollectionName)
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "serial", Value: 1},
		{Key: "type", Value: 1},
		{Key: "parent_serial", Value: 1},
		{Key: "geo_location", Value: 1},
	}
	opts := options.FindOne().SetProjection(projection)
	err := locationCollection.FindOne(ctx, bson.D{{Key: "serial", Value: serial}}, opts).Decode(&location)
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

func SearchLocations(ctx context.Context, query locationSchema.SearchLocation) (*responses.ListingResponse, error) {
	var locations []locationSchema.LocationGeneral = []locationSchema.LocationGeneral{}
	var listingResponse responses.ListingResponse
	listingResponse.TotalCount = 0
	listingResponse.List = locations
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), locationSchema.LocationCollectionName)
	limit := int64(query.Limit)
	page := int64(query.Page)
	sort := bson.D{{Key: "name", Value: 1}}
	skip := int64(page*limit - limit)
	match := bson.D{}
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "serial", Value: 1},
		{Key: "type", Value: 1},
		{Key: "parent_serial", Value: 1},
		{Key: "geo_location", Value: 1},
	}
	if len(query.Name) > 0 {
		names := strings.Split(query.Name, ",")
		match = append(match, bson.E{Key: "name", Value: bson.D{{Key: "$in", Value: names}}})
	}
	if len(query.Type) > 0 {
		types := strings.Split(query.Type, ",")
		match = append(match, bson.E{Key: "type", Value: bson.D{{Key: "$in", Value: types}}})
	}
	if query.ParentSerial > 0 {
		match = append(match, bson.E{Key: "parent_serial", Value: query.ParentSerial})
	}
	if len(query.SortBy) > 0 && query.SortOrder >= -1 {
		sort = bson.D{{Key: query.SortBy, Value: query.SortOrder}}
	}
	opts := options.Find().SetProjection(projection).SetSort(sort).SetSkip(skip).SetLimit(limit)
	results, err := automobileAdCollection.Find(ctx, match, opts)
	if err != nil {
		return nil, err
	}
	totalCount, err := automobileAdCollection.CountDocuments(ctx, match)
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
	listingResponse.TotalCount = uint64(totalCount)
	listingResponse.List = locations
	return &listingResponse, nil
}
