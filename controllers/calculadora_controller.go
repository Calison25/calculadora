package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Sum(context echo.Context) error {
	return context.JSON(http.StatusOK, "Funcionou Soma")
}

func Sub(context echo.Context) error {
	return context.JSON(http.StatusOK, "Funcionou substração")
}
