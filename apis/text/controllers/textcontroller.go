package controllers

import (
	"math/rand"

	"github.com/FournyP/kubernetes-tp/apis/text/ent"
	"github.com/FournyP/kubernetes-tp/apis/text/ent/text"
	"github.com/gofiber/fiber/v2"
)

type TextController struct {
	client *ent.Client
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewTextController(client *ent.Client) *TextController {
	return &TextController{
		client: client,
	}
}

func (controller *TextController) GetText(c *fiber.Ctx) error {
	name := c.Query("name", "null")

	text, err := controller.client.Text.Query().Where(text.NameEQ(name)).First(c.Context())

	if err != nil && !ent.IsNotFound(err) {
		panic(err)
	}

	if text == nil {
		textValue := randString(30)

		err = controller.client.Text.Create().
			SetText(textValue).
			SetName(name).
			Exec(c.Context())

		if err != nil {
			panic(err)
		}

		return c.JSON(textValue)
	}

	return c.JSON(text.Text)
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
