package controller

import (
	"graphyy/controller/scooter"
	"graphyy/repository"
)

// Controllers contains all the controllers
type Controllers struct {
	scooterController *scooter.Controller
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		scooterController: scooter.InitController(repositories.ScooterRepo),
	}
}
