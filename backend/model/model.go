package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Version          string `json:"version"`
	Port             int    `json:"port"`
	ServerPath       string `json:"serverPath"`
	Hostname         string `json:"hostname"`
	CaptchaSecretKey string `json:"captchaSecretKey"`
}

func ParseConfig(path string) (config *Config, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var ret Config
	json.Unmarshal(byteValue, &ret)
	return &ret, nil
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
