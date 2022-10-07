package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	InitEnv()
	client := InitOrm()
	defer client.Close()

	app := fiber.New()

	InitRoutes(app, client)

	app.Listen(":3000")
}
