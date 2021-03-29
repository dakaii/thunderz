package scooterrepo

import (
	"context"
	"fmt"
	"graphyy/internal/envvar"
	"graphyy/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ScooterRepo should i rename it?
type ScooterRepo struct {
	db *mongo.Database
}

// NewScooterRepo ..
func NewScooterRepo(db *mongo.Database) *ScooterRepo {
	return &ScooterRepo{
		db: db,
	}
}

func NewLocation(lat, long float64) model.Location {
	return model.Location{
		GeoJSONType: "Point",
		Coordinates: model.Coordinates{Latitude: lat, Longitude: long},
	}
}

//https://gist.github.com/Lebski/8f9b5992fec0bf175285f1c13b1e5051
// GetExistingUser fetches a user by the username from the db and returns it.
func (repo *ScooterRepo) GetScootersNearby(latitude float64, longitude float64, distance int) ([]model.Point, error) {
	var results []model.Point
	location := model.Point{Title: "scooter", Location: NewLocation(latitude, longitude)}
	pointCollection := repo.db.Collection(envvar.PointCollection())
	filter := bson.D{
		{Key: "location", Value: bson.D{
			{Key: "$near", Value: bson.D{
				{Key: "$geometry", Value: location},
				{Key: "$maxDistance", Value: distance},
			}},
		}},
	}
	cur, err := pointCollection.Find(context.Background(), filter)

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
