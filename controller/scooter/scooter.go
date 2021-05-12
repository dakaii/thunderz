package scooter

import (
	"graphyy/model"
	"graphyy/repository"
)

// declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type service interface {
	GetScootersNearby(lat float64, lng float64, distance int64, limit int64) ([]model.Point, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type Controller struct {
	service
}

// InitController initializes the user controller.
func InitController(scooterRepo *repository.ScooterRepo) *Controller {
	return &Controller{
		service: scooterRepo,
	}
}

func (c *Controller) GetNearbyScooters(lat float64, lng float64, distance int64, limit int64) ([]model.Point, error) {
	scooters, err := c.service.GetScootersNearby(lat, lng, distance, limit)
	return scooters, err
}
