package authentication

import (
	"github.com/labstack/echo"
)

// Register a new user
func Route(r *echo.Group, s Service) {
	// Routes
	r.POST("/authentication", s.GenerateCustomer)
	//only use with algorithm RS256
	r.POST("/validate", s.ValidateToken)

}
