package database

import (
	"context"
	"log"
	"time"

	"graphyy/internal/envvar"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// GetDatabase returns a database instance.
func InitDatabase() *mongo.Database {
	url := envvar.MongoURL()
	dbName := envvar.DBName()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(dbName)
	indexOpts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	pointIndexModel := mongo.IndexModel{
		Options: options.Index().SetBackground(true),
		Keys:    bsonx.MDoc{"location": bsonx.String("2dsphere")},
	}
	pointIndexes := db.Collection(envvar.PointCollection()).Indexes()
	_, err = pointIndexes.CreateOne(
		ctx,
		pointIndexModel,
		indexOpts,
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
