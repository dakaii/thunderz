package repository

import (
	"coldhongdae/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetExistingUser(username string) model.User
	SaveUser(user model.User) (model.User, error)
}

type UserRepo struct {
	db         *mongo.Database
	ctx        context.Context
	collection *mongo.Collection
}

// NewUserRepo ..
func NewUserRepo(db *mongo.Database, ctx context.Context, collection *mongo.Collection) *UserRepo {
	return &UserRepo{
		db:         db,
		ctx:        ctx,
		collection: collection,
	}
}

// GetExistingUser fetches a user by the username from the db and returns it.
func (h *UserRepo) GetExistingUser(username string) model.User {
	filter := bson.M{"username": username}
	var user model.User
	collection := h.collection
	ctx := h.ctx
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return model.User{Username: user.Username, Password: user.Password}
}

// SaveUser creates a new user in the db..
func (h *UserRepo) SaveUser(user model.User) (model.User, error) {
	// TODO handle the potential error below.
	hashedPass, _ := hashPassword(user.Password)
	user.Password = hashedPass

	collection := h.collection
	ctx := h.ctx
	insertResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a user with ID:", insertResult.InsertedID)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return model.User{Username: "", Password: ""}, nil
	}
	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}