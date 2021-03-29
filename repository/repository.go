package repository

import (
	"graphyy/repository/scooterrepo"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repositories contains all the repo structs
type Repositories struct {
	ScooterRepo *scooterrepo.ScooterRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *mongo.Database) *Repositories {
	scooterRepo := scooterrepo.NewScooterRepo(db)
	return &Repositories{ScooterRepo: scooterRepo}
}
