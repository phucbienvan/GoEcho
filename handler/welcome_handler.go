package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Welcome(c echo.Context)error  {
	return c.HTML(http.StatusOK, "<h1>Welcome to may app</h1>")
}
