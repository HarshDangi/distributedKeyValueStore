package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/harshdangi/distributedKeyValueStore/handlers"
	"github.com/harshdangi/distributedKeyValueStore/redisClient"
)

func getServerApp() *fiber.App {
	app := fiber.New()

	handlers.AttachHandlers(app)

	return app
}

func main() {
	if redisClient := redisClient.InitializeClient(); redisClient == nil {
		log.Fatal("Unable to connect to Redis.")
		return
	}
	app := getServerApp()

	app.Listen(":3000")
}
