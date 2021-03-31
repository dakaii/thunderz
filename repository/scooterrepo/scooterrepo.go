package scooterrepo

import (
	"context"
	"fmt"
	"graphyy/internal/envvar"
	"graphyy/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ScooterRepo should i rename it?
type ScooterRepo struct {
	db *mongo.Database
}

// NewScooterRepo constructs a ScooterRepo
func NewScooterRepo(db *mongo.Database) *ScooterRepo {
	return &ScooterRepo{
		db: db,
	}
}

//https://gist.github.com/Lebski/8f9b5992fec0bf175285f1c13b1e5051
// GetScootersNearby fetches the scooters within the specified distance.
func (repo *ScooterRepo) GetScootersNearby(latitude float64, longitude float64, distance int64, limit int64) ([]model.Point, error) {
	var results []model.Point
	pointCollection := repo.db.Collection(envvar.PointCollection())
	filter := bson.D{
		{Key: "location", Value: bson.D{
			{Key: "$near", Value: bson.D{
				{Key: "$geometry", Value: model.Location{
					GeoJSONType: "Point",
					Coordinates: []float64{longitude, latitude}},
				},
				{Key: "$maxDistance", Value: distance},
			}},
		}},
	}
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	cur, err := pointCollection.Find(context.Background(), filter, findOptions)

	if err != nil {
		return []model.Point{}, err
	}

	for cur.Next(context.TODO()) {
		var elem model.Point
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Could not decode Point")
			return []model.Point{}, err
		}
		results = append(results, elem)
	}

	return results, nil
}
