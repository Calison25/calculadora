package routes

import (
	"calculadora/controllers"
	"github.com/labstack/echo"
	"net/http"
)

func GetRoutes(e *echo.Echo) {
	e.GET("/hello", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Hello World")
	})

	e.POST("/sum", controllers.Sum)
	e.POST("/sub", controllers.Sub)

	e.Logger.Fatal(e.Start(":1500"))
}
