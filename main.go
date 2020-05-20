package main

import (
	"github.com/janakhpon/gopherscom/src/config"

	"os"

	"github.com/janakhpon/gopherscom/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	config.Connect()
	config.ConnectRedis()
	godotenv.Load()
	port := os.Getenv("PORT")
	mode := os.Getenv("MODE")
	router := routes.ExtRouter(mode)
	router.Run(":" + port)
}
