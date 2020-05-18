package main

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/saul8888/postgrsql/postgrsql/authentication"
	"github.com/saul8888/postgrsql/postgrsql/dbpostgrsql"
	"github.com/saul8888/postgrsql/postgrsql/employee"
	"github.com/saul8888/postgrsql/postgrsql/middelware"
)

func main() {
	db := dbpostgrsql.ConnectDB()
	route := echo.New()
	r := route.Group("/api")
	//----------------------------//
	r.GET("/", hello)
	r.GET("/prueba", test)
	//----------------------------//
	//Database
	var data = dbpostgrsql.NewDataBase(db)
	//authentication
	var autheService = authentication.NewService(data)
	//employee
	var Semployee = employee.NewService(data)

	//----------------------------------------------//
	authentication.Route(r, autheService)
	middelware.ConfigMiddelware(r)
	//----------------------------------------------//
	employee.Route(r, Semployee)

	// Start the server on localhost port 8080 and log any errors
	route.Logger.Fatal(route.Start(":8080"))

}

//----------------test----------------------//
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "POSTGRSQL")
}

func test(c echo.Context) error {
	example := &dbpostgrsql.Employee{
		Name:     "test",
		Email:    "test@example.com",
		Password: "test1234",
	}
	db1 := dbpostgrsql.ConnectDB()
	var data1 = dbpostgrsql.NewDataBase(db1)
	pr, err := data1.Insert(example)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, pr)
}
