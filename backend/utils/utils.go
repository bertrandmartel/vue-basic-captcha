package utils

import (
	"net/http"

	"github.com/bertrandmartel/vue-basic-captcha/backend/model"
	"github.com/labstack/echo/v4"
)

func ProcessResult(i interface{}, err error, c echo.Context) error {
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
			Error:            "server_error",
			ErrorDescription: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, i)
}

func SendError(errorMessage string, errorDescription string) *model.ErrorResponse {
	return &model.ErrorResponse{
		Error:            errorMessage,
		ErrorDescription: errorDescription,
	}
}
