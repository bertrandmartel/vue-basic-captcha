package middleware

import (
	"net/http"

	"github.com/bertrandmartel/vue-basic-captcha/backend/application"
	"github.com/bertrandmartel/vue-basic-captcha/backend/session"
	"github.com/bertrandmartel/vue-basic-captcha/backend/utils"
	"github.com/labstack/echo/v4"
)

func UseSession(next echo.HandlerFunc, context echo.Context, app application.BasicCaptchaApp, api bool) error {
	cookie, err := app.GetCookie(context, session.SessionCookie)
	if err != nil {
		if !api {
			return context.Redirect(http.StatusFound, "/login")
		}
		return context.JSON(http.StatusUnauthorized, utils.SendError("unauthorized", "authentication failed"))
	}
	s, err := app.GetSessionFromStore(&cookie)
	if err != nil {
		return err
	}
	app.SetSessionContext(context, s)
	return next(context.(echo.Context))
}

func UseClearSession(context echo.Context, app application.BasicCaptchaApp) {
	cookie, err := app.GetCookie(context, session.SessionCookie)
	if err == nil {
		app.DeleteSession(&cookie)
	}
	app.DeleteCookie(context, session.SessionCookie)
}
