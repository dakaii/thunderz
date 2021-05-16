package repository

import "graphyy/storage"

// Repositories contains all the repo structs
type Repositories struct {
	ScooterRepo *ScooterRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db storage.Storage) *Repositories {
	scooterRepo := NewScooterRepo(db)
	return &Repositories{ScooterRepo: scooterRepo}
}
