package main

import (
	"github.com/FournyP/kubernetes-tp/apis/image/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	pingController := controllers.NewPingController()
	imageController := controllers.NewImageController()

	app.Get("/ping", pingController.Ping)
	app.Get("/image", imageController.GetImage)
}
