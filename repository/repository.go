package repository

import (
	"graphyy/database"
)

// Repositories contains all the repo structs
type Repositories struct {
	ScooterRepo *ScooterRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db database.Storage) *Repositories {
	scooterRepo := NewScooterRepo(db)
	return &Repositories{ScooterRepo: scooterRepo}
}
