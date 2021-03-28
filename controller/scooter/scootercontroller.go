package scooter

import (
	"graphyy/model"
	"graphyy/repository/scooterrepo"
)

// declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type repository interface {
	GetScootersNearby(latitude string, longitude string) model.Scooter
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
