package main

import (
	"github.com/FournyP/kubernetes-tp/apis/text/controllers"
	"github.com/FournyP/kubernetes-tp/apis/text/ent"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, client *ent.Client) {
	pingController := controllers.NewPingController()
	textController := controllers.NewTextController(client)

	app.Get("/ping", pingController.Ping)
	app.Get("/text", textController.GetText)
}
