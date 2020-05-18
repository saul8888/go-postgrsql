package main

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/saul8888/postgrsql/postgrsql/dbpostgrsql"
)

func main() {
	///db := dbpostgrsql.ConnectDB()
	route := echo.New()
	r := route.Group("/api")
	//----------------------------//
	r.GET("/", hello)
	r.GET("/prueba", prueba)

	//Database
	///var data = dbpostgrsql.NewDataBase(db)
	//authentication
	///var autheService = authentication.NewService(data)
	//customer
	///var people = employee.NewService(data)

	//----------------------------------------------//
	///authentication.Route(r, autheService)

	///middelware.ConfigMiddelware(r)
	///employee.Route(r, people)

	// Start the server on localhost port 8000 and log any errors
	route.Logger.Fatal(route.Start(":8080"))

}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "POSTGRSQL")
}

func prueba(c echo.Context) error {
	db1 := dbpostgrsql.ConnectDB()
	var data1 = dbpostgrsql.NewDataBase(db1)
	pr, err := data1.GetCount()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, pr)
}
