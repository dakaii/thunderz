package database

import (
	"context"
	"graphyy/internal"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Storage struct {
	db *mongo.Database
}

// GetDatabase returns a database instance.
func InitDatabase() Storage {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(internal.MongoURL))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(internal.DBName)
	storage := Storage{db}
	collection := storage.PointerCollection()
	models := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "location", Value: bsonx.String("2dsphere")}},
		},
	}

	// Declare an options object
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err = collection.Indexes().CreateMany(ctx, models, opts)
	if err != nil {
		log.Fatal(err)
	}
	return storage
}

func (storage *Storage) PointerCollection() *mongo.Collection {
	return storage.db.Collection("scooter")
}
