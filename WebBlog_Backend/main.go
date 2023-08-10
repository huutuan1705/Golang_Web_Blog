package main

import (
	"example.com/blog/database"
	"example.com/blog/routes"
	"github.com/gofiber/fiber/v2"
	//"net/http"
)

func main() {
	database.Connect()

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8080")
}
