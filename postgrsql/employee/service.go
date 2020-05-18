package employee

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/saul8888/postgrsql/postgrsql/dbpostgrsql"
)

type Service interface {
	GetEmployeeById(context echo.Context) error
	GetEmployees(context echo.Context) error
	InsertEmployee(context echo.Context) error
	UpdateEmployee(context echo.Context) error
	DeleteEmployee(context echo.Context) error
}

type service struct {
	repo dbpostgrsql.Postgrsql
}

func NewService(repo dbpostgrsql.Postgrsql) Service {
	return &service{repo: repo}
}

func (s *service) GetEmployeeById(c echo.Context) error {
	employeeID := c.QueryParam("ID")
	row, err := s.repo.GetById(employeeID)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, row)
}

func (s *service) InsertEmployee(c echo.Context) (err error) {
	employee := new(dbpostgrsql.Employee)
	if err = c.Bind(employee); err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	newEmployee, err := s.repo.Insert(employee)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, newEmployee)
}

func (s *service) GetEmployees(c echo.Context) (err error) {
	params := new(dbpostgrsql.GetEmployeesRequest)
	if err = c.Bind(params); err != nil {
		return
	}
	row, err := s.repo.GetTotal(params)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	totalEmployees, err := s.repo.GetCount()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	cant := dbpostgrsql.EmployeeList{Data: row, TotalRecords: totalEmployees}
	return c.JSON(http.StatusOK, cant)

}

func (s *service) UpdateEmployee(c echo.Context) (err error) {
	employee := new(dbpostgrsql.UpdateEmployeeRequest)
	if err = c.Bind(employee); err != nil {
		return
	}
	update, err := s.repo.Update(employee)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	//result := StatusAction{Action: "OK", Update: updateEmployee}
	return c.JSON(http.StatusOK, update)
}

func (s *service) DeleteEmployee(c echo.Context) error {
	employeeID := c.QueryParam("ID")
	delete, err := s.repo.DeleteById(employeeID)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, delete)
}
