package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ConfiguratonResponse struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func Configuration(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, &ConfiguratonResponse{
		Topic:   "Welcome",
		Message: "This is the result of an authenticated API",
	})
}
