package controllers

import "github.com/gofiber/fiber/v2"

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (controller *PingController) Ping(c *fiber.Ctx) error {
	return c.JSON("PONG")
}
