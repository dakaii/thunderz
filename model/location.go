package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	GeoJSONType string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Point struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title"`
	Location Location           `json:"location"`
}
