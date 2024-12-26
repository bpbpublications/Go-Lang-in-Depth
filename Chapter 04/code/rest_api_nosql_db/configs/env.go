package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {

	fmt.Println("Getting from Env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}

	return os.Getenv("MONGOURI")
}
