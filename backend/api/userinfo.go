package api

import (
	"net/http"

	"github.com/bertrandmartel/vue-basic-captcha/backend/application"
	"github.com/bertrandmartel/vue-basic-captcha/backend/session"
	"github.com/labstack/echo/v4"
)

type UserInfoResponse struct {
	Authenticated bool `json:"authenticated"`
}

func UserInfo(c echo.Context) (err error) {
	app := *c.Get("authApp").(*application.BasicCaptchaApp)
	_, err = app.GetCookie(c, session.SessionCookie)
	if err != nil {
		return c.JSON(http.StatusOK, &UserInfoResponse{
			Authenticated: false,
		})
	}
	return c.JSON(http.StatusOK, &UserInfoResponse{
		Authenticated: true,
	})
}
