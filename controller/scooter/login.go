package scooter

import (
	"graphyy/model"
)

// Login returns a jwt.
func (c *Controller) GetNearbyScooters(latitude float64, longitude float64, distance int) ([]model.Point, error) {
	scooters, err := c.service.GetScootersNearby(latitude, longitude, distance)
	// if existingUser.Username == "" {
	// 	return model.AuthToken{}, errors.New("No user found with the inputted username")
	// }
	// isValid := checkPasswordHash(user.Password, existingUser.Password)
	// if !isValid {
	// 	return model.AuthToken{}, errors.New("Invalid Credentials")
	// }

	return scooters, err
}
