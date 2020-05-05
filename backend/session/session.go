package session

import (
	"net/http"
)

const SessionCookie = "BASIC-CAPTCHA"

type Session struct {
	ID         string `json:"id"`
	HTTPClient *http.Client
}
