package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/bertrandmartel/vue-basic-captcha/backend/application"
	"github.com/bertrandmartel/vue-basic-captcha/backend/model"

	"github.com/bertrandmartel/vue-basic-captcha/backend/api"
	authMiddleware "github.com/bertrandmartel/vue-basic-captcha/backend/middleware"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	config, err := model.ParseConfig("config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("[basic-captcha] version %v\n", config.Version)
	log.Printf("[basic-captcha] server path %v:%v\n", config.ServerPath, config.Port)

	hostnameEnv := getenv("HOSTNAME", "")
	capchaSecretKeyEnv := getenv("CAPTCHA_SECRET_KEY", "")
	if hostnameEnv != "" {
		config.Hostname = hostnameEnv
	}
	if capchaSecretKeyEnv != "" {
		config.CaptchaSecretKey = capchaSecretKeyEnv
	}

	redisHostEnv := getenv("REDIS_HOST", "localhost")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHostEnv + ":6379",
		Password: "",
		DB:       1,
	})
	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	e := echo.New()
	/*
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		}))
	*/
	UseCommonMiddleware(e)
	routes(e, config, redisClient, httpClient)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Port)))
}

func routes(e *echo.Echo, cfg *model.Config, conn *redis.Client, httpClient *http.Client) {
	var authHandler application.BasicCaptchaApp = &application.AuthApp{
		Config:      cfg,
		RedisClient: conn,
		HTTPClient:  httpClient,
	}
	e.Use(bindApp(&authHandler))

	e.POST("/login", api.Login)
	e.GET("/userinfo", api.UserInfo)
	e.GET("/logout", func(c echo.Context) error {
		return c.(echo.Context).Redirect(http.StatusFound, "/login")
	}, MWClearSession)

	//API that require authentication
	e.GET("/configuration", api.Configuration, MWApiSession)

	if os.Getenv("APP_ENV") == "development" {
		devURL, err := url.Parse("http://localhost:8080")
		if err != nil {
			e.Logger.Fatal(err)
		}
		targets := []*middleware.ProxyTarget{
			{
				URL: devURL,
			},
		}
		e.GET("/login", func(c echo.Context) error { return nil }, middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

		//Views that require authentication
		e.Group("/", MWSession, middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

		//Views without authentication
		e.Group("/*", middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))
	} else {
		e.File("/login", "dist/index.html")

		//Views that require authentication
		e.File("/", "dist/index.html", MWSession)

		//Views without authentication
		e.Group("/*", middleware.StaticWithConfig(middleware.StaticConfig{
			Root:   "dist",
			Browse: false,
		}))
	}
}

func bindApp(app *application.BasicCaptchaApp) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("authApp", app)
			return next(c)
		}
	}
}

//middleware for validation
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func MWClearSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		app := c.Get("authApp").(*application.BasicCaptchaApp)
		authMiddleware.UseClearSession(c, *app)
		return next(c)
	}
}

func MWSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		app := c.Get("authApp").(*application.BasicCaptchaApp)
		return authMiddleware.UseSession(next, c, *app, false)
	}
}
func MWApiSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		app := c.Get("authApp").(*application.BasicCaptchaApp)
		return authMiddleware.UseSession(next, c, *app, true)
	}
}

func UseCommonMiddleware(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip} - - ${time_rfc3339_nano} \"${method} ${uri} ${protocol}\" ${status} ${bytes_out} \"${referer}\" \"${user_agent}\"\n",
	}))
	e.Use(middleware.Recover())
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
