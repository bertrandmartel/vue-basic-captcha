package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/go-redis/redis/v7"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	redisHostEnv := getenv("REDIS_HOST", "localhost")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHostEnv + ":6379",
		Password: "",
		DB:       1,
	})
	log.Println("Enter a username")
	username := getField()

	//password
	fmt.Print("Enter password: ")
	bytePassword, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Print("\n")
	password := string(bytePassword)
	if password == "" {
		log.Println("password is required")
		return
	}

	hash := hashAndSalt(bytePassword)

	user := string(username)
	err = setCredentials(&user, &hash, redisClient)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("check user")
	entry, err := getCredentials(&user, redisClient)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("found : " + *entry)
}

func getField() []byte {
	var field string
	_, err := fmt.Scan(&field)
	if err != nil {
		log.Println(err)
	}
	return []byte(field)
}

//https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func setCredentials(username *string, password *string, conn *redis.Client) error {
	err := conn.Set("user:"+*username, *password, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func getCredentials(username *string, conn *redis.Client) (password *string, err error) {
	r, err := conn.Get("user:" + *username).Result()
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
