package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context)error  {
	return c.JSON(http.StatusOK, echo.Map{
		"user" : "phuc",
		"email" : "phuc@gmail",
	})
}
