package main

import (
	"jwt-token/repositories"
	"jwt-token/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	repositories.Connect()
	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))

}
