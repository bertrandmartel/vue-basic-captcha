package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/bertrandmartel/vue-basic-captcha/backend/application"
	"github.com/bertrandmartel/vue-basic-captcha/backend/session"
	"github.com/bertrandmartel/vue-basic-captcha/backend/utils"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Username          string `form:"login"`
	Password          string `form:"password"`
	RecaptchaResponse string `form:"g-recaptcha-response"`
}

type CaptchaResponse struct {
	Success    bool     `json:"success"`
	Timestamp  string   `json:"challenge_ts"`
	Hostname   string   `json:"hostname"`
	ErrorCodes []string `json:"error-codes"`
}

func Login(c echo.Context) (err error) {
	app := *c.Get("authApp").(*application.BasicCaptchaApp)
	data := new(LoginResponse)
	if err = c.Bind(data); err != nil {
		return
	}
	captchaResponse := new(CaptchaResponse)
	err = verifyCaptcha(app.GetHTTPClient(), app.GetConfig().CaptchaSecretKey, data.RecaptchaResponse, c.RealIP(), captchaResponse)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.SendError("unauthorized", "authentication failed"))
	}
	log.Println(captchaResponse)
	if !captchaResponse.Success {
		return c.JSON(http.StatusUnauthorized, utils.SendError("unauthorized", "authentication failed"))
	}
	if captchaResponse.Hostname != app.GetConfig().Hostname {
		return c.JSON(http.StatusUnauthorized, utils.SendError("unauthorized", "authentication failed"))
	}
	if data.Username == "" && data.Password == "" {
		return c.JSON(http.StatusUnauthorized, utils.SendError("unauthorized", "authentication failed"))
	}
	hash, err := app.GetHashFromUsername(&data.Username)
	if err != nil {
		log.Println("user not found")
		return c.JSON(http.StatusUnauthorized, utils.SendError("unauthorized", "authentication failed"))
	}
	result := comparePasswords(*hash, []byte(data.Password))
	if !result {
		log.Println("invalid password")
		return c.JSON(http.StatusUnauthorized, utils.SendError("unauthorized", "authentication failed"))
	}
	s := &session.Session{
		ID: uuid.NewV4().String(),
	}
	sessionVal, err := app.SetSession(s)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, utils.SendError("unauthorized", "authentication failed"))
	}
	app.SetSessionCookie(c, session.SessionCookie, sessionVal)
	app.SetSessionContext(c, s)
	return c.NoContent(http.StatusOK)
}

func verifyCaptcha(httpClient *http.Client, captchaSecretKey string, captchaResponse string, remoteip string, target interface{}) error {
	if httpClient == nil {
		return errors.New("no http client specified")
	}
	form := url.Values{}
	form.Add("secret", captchaSecretKey)
	form.Add("response", captchaResponse)
	form.Add("remoteip", remoteip)
	req, err := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode == 404 {
		return errors.New("record was not found")
	}
	if r.StatusCode != 200 {
		return errors.New("received incorrect status : " + strconv.Itoa(r.StatusCode))
	}
	return json.NewDecoder(r.Body).Decode(target)
}

//https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
