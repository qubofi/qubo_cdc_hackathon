package main

import (
	"os"
	"qubo/qubo-backend/infrastructure/database"
	"qubo/qubo-backend/infrastructure/redis"
	"qubo/qubo-backend/interfaces/router"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {
	redis.SetupRedis()
	database.SetupDatabase()
	router.SetupRouter().Run(":8080")
}
