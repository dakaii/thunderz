package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func main() {
	conn := initDatabase()
	for i := 0; i < 150; i++ {
		lat := randomFloat(1.48, 1.21)
		lon := randomFloat(104.1, 103.5)
		point := Point{Title: "scooter", Location: Location{
			GeoJSONType: "Point",
			Coordinates: Coordinates{Latitude: lat, Longitude: lon},
		}}
		addPoint(conn, point)
	}
}

func randomFloat(max float64, min float64) float64 {
	return min + rand.Float64()*(max-min)
}
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("config.conf")

	if err != nil {
		log.Fatalf("Error loading config.conf file")
	}

	return os.Getenv(key)
}

func initDatabase() *mongo.Client {
	url := goDotEnvVariable("MONGODB_URL")
	fmt.Printf("Mongo URL: %s\n", url)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func addPoint(conn *mongo.Client, point Point) error {
	dbName := goDotEnvVariable("MONGODB_DB_NAME")
	collectionName := goDotEnvVariable("MONGODB_COLLECTION_SCOOTER")

	coll := conn.Database(dbName).Collection(collectionName)
	point.ID = primitive.NewObjectID()
	point.Title = "scooter"
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	fmt.Printf("Could not insert new Point. Id: %s\n", point.ID)
	fmt.Printf("Could not insert new Point. Id: %s\n", point.Location.Coordinates.Latitude)
	fmt.Printf("Could not insert new Point. Id: %s\n", point.Location.Coordinates.Longitude)
	insertResult, err := coll.InsertOne(ctx, point)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Printf("Could not insert new Point. Id: %s\n", point.ID)
		return err
	}
	fmt.Printf("Inserted new Point. ID: %s\n", insertResult.InsertedID)
	return nil
}
