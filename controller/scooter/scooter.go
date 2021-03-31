package scooter

import (
	"graphyy/model"
	"graphyy/repository/scooterrepo"
)

// declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type repository interface {
	GetScootersNearby(latitude float64, longitude float64, distance int64, limit int64) ([]model.Point, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type Controller struct {
	service repository
}

// InitController initializes the user controller.
func InitController(scooterRepo *scooterrepo.ScooterRepo) *Controller {
	return &Controller{
		service: scooterRepo,
	}
}

func (c *Controller) GetNearbyScooters(latitude float64, longitude float64, distance int64, limit int64) ([]model.Point, error) {
	scooters, err := c.service.GetScootersNearby(latitude, longitude, distance, limit)
	return scooters, err
}
