package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v7"
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
	user := string(username)
	err := redisClient.Del("user:" + user).Err()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("user " + user + " was deleted")
}

func getField() []byte {
	var field string
	_, err := fmt.Scan(&field)
	if err != nil {
		log.Println(err)
	}
	return []byte(field)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
