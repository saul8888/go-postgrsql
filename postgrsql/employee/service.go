package employee

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/saul8888/postgrsql/postgrsql/dbpostgrsql"
)

type Service interface {
	GetCustomerById(context echo.Context) error
	GetCustomers(context echo.Context) error
	InsertCustomer(context echo.Context) error
	UpdateCustomer(context echo.Context) error
	DeleteCustomer(context echo.Context) error
}

type service struct {
	repo dbpostgrsql.Postgrsql
}

func NewService(repo dbpostgrsql.Postgrsql) Service {
	return &service{repo: repo}
}

func (s *service) GetCustomerById(c echo.Context) error {
	customerID := c.QueryParam("ID")
	row, err := s.repo.GetById(customerID)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, row)
}

func (s *service) InsertCustomer(c echo.Context) (err error) {
	customer := new(dbpostgrsql.Customer)
	if err = c.Bind(customer); err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	newCustomer, err := s.repo.Insert(customer)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, newCustomer)
}

func (s *service) GetCustomers(c echo.Context) (err error) {
	params := new(dbpostgrsql.GetCustomersRequest)
	if err = c.Bind(params); err != nil {
		return
	}
	row, err := s.repo.GetTotal(params)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	totalCustomers, err := s.repo.GetCount()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	cant := dbpostgrsql.CustomerList{Data: row, TotalRecords: totalCustomers}
	return c.JSON(http.StatusOK, cant)

}

func (s *service) UpdateCustomer(c echo.Context) (err error) {
	customer := new(dbpostgrsql.UpdateCustomerRequest)
	if err = c.Bind(customer); err != nil {
		return
	}
	update, err := s.repo.Update(customer)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	//result := StatusAction{Action: "OK", Update: updateCustomer}
	return c.JSON(http.StatusOK, update)
}

func (s *service) DeleteCustomer(c echo.Context) error {
	customerID := c.QueryParam("ID")
	delete, err := s.repo.DeleteById(customerID)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, delete)
}
