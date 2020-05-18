package employee

import (
	"github.com/labstack/echo"
)

// Register a new user
func Route(r *echo.Group, s Service) {
	// Routes
	r.GET("/getId", s.GetCustomerById)
	r.GET("/total", s.GetCustomers)
	r.POST("/insert", s.InsertCustomer)
	r.PUT("/update", s.UpdateCustomer)
	r.DELETE("/delete", s.DeleteCustomer)

}
