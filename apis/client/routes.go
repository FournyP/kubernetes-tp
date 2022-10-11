package main

import (
	"github.com/FournyP/kubernetes-tp/apis/client/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	pingController := controllers.NewPingController()
	containerController := controllers.NewContainerController()

	app.Get("/ping", pingController.Ping)
	app.Post("/create/container/:name", containerController.CreateContainer)
}
