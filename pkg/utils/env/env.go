package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	load()
}

func load() {
	file := "dev/.env"
	err := godotenv.Load(file)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Get(key string) string {
	val, ok := os.LookupEnv(key)

	if ok {
		return val
	} else {
		log.Panicf("Error finding in .env file")
		return ""
	}
}
