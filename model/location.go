package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Location struct {
	GeoJSONType string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Point struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title" bson:""`
	Location Location           `json:"location" bson:"type"`
}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"createdAt" graphql:"-"`
	UpdatedAt time.Time          `json:"updatedAt" graphql:"-"`
	DeletedAt gorm.DeletedAt     `json:"deletedAt" graphql:"-"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
}

//AuthToken struct
type AuthToken struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"accessToken"`
	ExpiresIn int64  `json:"expiresIn"`
}

//AuthTokenClaim struct
type AuthTokenClaim struct {
	jwt.StandardClaims
	User
}
