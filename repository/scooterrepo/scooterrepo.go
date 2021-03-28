package scooterrepo

import (
	"context"
	"graphyy/model"

	"go.mongodb.org/mongo-driver/mongo"
)

// ScooterRepo should i rename it?
type ScooterRepo struct {
	ctx context.Context
	db  *mongo.Database
}

// NewScooterRepo ..
func NewScooterRepo(ctx context.Context, db *mongo.Database) *ScooterRepo {
	return &ScooterRepo{
		ctx: ctx,
		db:  db,
	}
}

// GetExistingUser fetches a user by the username from the db and returns it.
func (repo *ScooterRepo) GetScootersNearby(latitude string, longitude string) model.Scooter {
	// filter := bson.M{"username": username}
	// var user model.User
	// err := h.collection.FindOne(h.ctx, filter).Decode(&user)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// return model.User{Username: user.Username, Password: user.Password}
	return model.Scooter{Latitude: latitude, Longitude: longitude}
}
