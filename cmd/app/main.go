package main

import (
	"log"
	"os"

	"github.com/chalermridp/go-template-rest/internal/app/config"
	"github.com/chalermridp/go-template-rest/internal/app/router"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error occurred during load .env")
	}
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.NewRouter(init)

	err := app.Run(":" + port)
	if err != nil {
		return
	}
}
