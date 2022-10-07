package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	InitRoutes(app)

	app.Listen(":3000")
}
