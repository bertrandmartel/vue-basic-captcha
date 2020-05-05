package application

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bertrandmartel/vue-basic-captcha/backend/model"
	"github.com/bertrandmartel/vue-basic-captcha/backend/session"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
)

type BasicCaptchaApp interface {
	GetHTTPClient() *http.Client
	SetSession(*session.Session) (id string, e error)
	GetSessionFromStore(uuid *string) (s *session.Session, e error)
	DeleteSession(uuid *string) error
	GetConfig() *model.Config
	SetCookie(c echo.Context, name string, value string)
	GetCookie(c echo.Context, name string) (string, error)
	DeleteCookie(c echo.Context, name string)
	SetSessionCookie(c echo.Context, name string, value string)
	SetSessionContext(c echo.Context, s *session.Session)
	Redirect(c echo.Context, location *string) error
	GetHashFromUsername(username *string) (hash *string, e error)
}

type AuthApp struct {
	Config      *model.Config
	RedisClient *redis.Client
	HTTPClient  *http.Client
}

func (app *AuthApp) SetSession(session *session.Session) (id string, e error) {
	sessionKey := session.ID
	sessionJSON, err := json.Marshal(*session)
	if err != nil {
		return "", err
	}
	err = app.RedisClient.Set("session:"+sessionKey, sessionJSON, time.Hour).Err()
	if err != nil {
		return "", err
	}
	return sessionKey, nil
}
func (app *AuthApp) GetSessionFromStore(uuid *string) (s *session.Session, e error) {
	r, err := app.RedisClient.Get("session:" + *uuid).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(r), &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
func (app *AuthApp) DeleteSession(uuid *string) error {
	return app.RedisClient.Del(*uuid).Err()
}
func (app *AuthApp) SetCookie(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}
func (app *AuthApp) GetCookie(c echo.Context, name string) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
func (app *AuthApp) DeleteCookie(c echo.Context, name string) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return
	}
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

}
func (app *AuthApp) SetSessionCookie(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}
func (app *AuthApp) SetSessionContext(c echo.Context, s *session.Session) {
	c.Set("session", s)
}
func (app *AuthApp) GetConfig() *model.Config {
	return app.Config
}
func (app *AuthApp) Redirect(c echo.Context, location *string) error {
	return c.Redirect(http.StatusFound, *location)
}
func (app *AuthApp) GetHTTPClient() *http.Client {
	return app.HTTPClient
}
func (app *AuthApp) GetHashFromUsername(username *string) (hash *string, e error) {
	user, err := app.RedisClient.Get("user:" + *username).Result()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
