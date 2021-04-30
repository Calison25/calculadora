package main

import (
	"calculadora/routes"
	"github.com/labstack/echo"
)

func main()  {
	e := echo.New()
	routes.GetRoutes(e)
	e.Logger.Fatal(e.Start(":1500"))
}