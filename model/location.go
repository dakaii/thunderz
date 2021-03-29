package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct
type Coordinates struct {
	Latitude  float64 `json:"latidute"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
	GeoJSONType string      `json:"type" bson:"type"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates"`
}

type Point struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title"`
	Location Location           `json:"location"`
}
