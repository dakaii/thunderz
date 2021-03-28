package database

import (
	"context"
	"log"
	"time"

	"graphyy/internal/envvar"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetDatabase returns a database instance.
func InitDatabase() (context.Context, *mongo.Database) {
	url := envvar.MongoURL()
	// collectionNameScooter := constant.CollectionNameScooter()
	dbName := envvar.DBName()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(dbName)
	// collectionScooter := database.Collection(collectionNameScooter)
	// opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	// model := []mongo.IndexModel{
	// 	{
	// 		Keys: bsonx.Doc{{Key: "username", Value: bsonx.String("text")}},
	// 	},
	// }
	// _, err = collectionScooter.Indexes().CreateMany(ctx, model, opts)
	return ctx, database
}
