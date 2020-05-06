# Vue Basic Captcha

![build](https://github.com/bertrandmartel/vue-basic-captcha/workflows/build/badge.svg) [![License](http://img.shields.io/:license-mit-blue.svg)](LICENSE.md)

Template project of a Vue.js application with basic authentication & captcha integration in Go

This template is great for making a quick responsive Vue.js project protected by username/password with captcha with vuetify and vuex out of the box. Redis is used to store the session & the username/password (password is hashed & salted).

![vuejs](https://user-images.githubusercontent.com/5183022/81025085-bd82b880-8e75-11ea-9d45-814ece3f9697.png)

## Install Demo

#### Using binaries

[Download the latest release](https://github.com/bertrandmartel/vue-basic-captcha/releases/latest)

* start ngrok on port 6004
* update host & captcha secret key in config.json
* go to `http://your.host.ngrok.io:6004`

Using binaries, you must have redis installed 

#### Using Docker

```bash
git clone git://github.com/bertrandmartel/vue-basic-captcha.git
cd vue-basic-captcha
```

* start ngrok on port 6004 and note down the port in the HOSTNAME variable in docker-compose below

Create docker-compose.yml :

```yaml
version: '2'

services:
  backend:
    build: .
    environment:
      REDIS_HOST: redis
      HOSTNAME: your.host.com
      CAPTCHA_SECRET_KEY: [captcha secret key here]
  redis:
    image: redis:6.0
```

```bash
docker-compose up
```

Create a user using the provisioning tool :
```
docker exec -it $(docker ps | grep vue-basic-captcha_backend | cut -d ' ' -f1) sh
cd ../provisioning/add
./add
```

go to `http://your.host.ngrok.io:6004`

#### Developer mode

In developer mode, you must have redis installed 

```bash
git clone git://github.com/bertrandmartel/vue-basic-captcha.git
cd vue-basic-captcha
cd backend
make install
```

create `.env.secret` in backend directory with content : 

```bash
HOSTNAME=your.host.ngrok.io
CAPTCHA_SECRET_KEY=[captcha secret key here]
```

Then run the backend :

```basj
make run
```

in another tab

```
cd frontend
npm i
vue ui
```

go to `http://your.host.ngrok.io:6004`

## Provisioning tools

2 programs provide way to add & remove a user to/from Redis :

* add a user

```bash
cd ./provisioning
go run ./add
```

* delete a user

```bash
cd ./provisioning
go run ./delete
```

## Configuration

2 way of configuring this app :

* config.json

```
{
    "version": "0.1",
    "port": 6004,
    "serverPath": "http://localhost",
    "hostname": "your.host.com",
    "captchaSecretKey": "[captcha secret key here]"
}
```

* environment variables 

|  Name  | Sample values  | Description |
|--------|--------------|---------------|
| REDIS_HOST |  redis / localhost |  Redis hostname |
| HOSTNAME   |  your.host.com     | hostname to be checked in captcha response |
| CAPTCHA_SECRET_KEY |  [captchaKey]  | captcha secret key | 

## Open Source components

* Backend in Go

  * [echo](https://echo.labstack.com/)
  * [go-redis](https://github.com/go-redis/redis)

* Frontend in Javascript

  * [vue.js](https://vuejs.org/)
  * [vuetify](https://vuetifyjs.com/en/getting-started/quick-start/)
  * [vuex](https://vuex.vuejs.org/)
  * [vue-router](https://router.vuejs.org/)

## External resources

* [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) is used to hash/salt the password in provisioning tools and when comparigng hashes. [This post](https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72) was very helpful

* Login view inspired by [this post](https://medium.com/vue-mastery/getting-started-with-vuetify-2-0-522ad3a55154)