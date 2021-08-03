package main

import (
	"github.com/janakhpon/gopherscom/src/config"

	"os"

	"github.com/janakhpon/gopherscom/src/routes"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	_ = godotenv.Load(".env")
	config.Connect()
	config.ConnectRedis()
	port := os.Getenv("PORT")
	print("port ", port)
	mode := os.Getenv("MODE")
	router := routes.ExtRouter(mode)
	router.Run(":" + port)
}
