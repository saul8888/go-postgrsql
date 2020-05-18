package employee

import (
	"github.com/labstack/echo"
)

// Register a new user
func Route(r *echo.Group, s Service) {
	// Routes
	r.GET("/getId", s.GetEmployeeById)
	r.GET("/total", s.GetEmployees)
	r.POST("/insert", s.InsertEmployee)
	r.PUT("/update", s.UpdateEmployee)
	r.DELETE("/delete", s.DeleteEmployee)

}
