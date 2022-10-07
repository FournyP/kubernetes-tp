package controllers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type ImageController struct{}

func NewImageController() *ImageController {
	return &ImageController{}
}

func (controller *ImageController) GetImage(c *fiber.Ctx) error {
	name := c.Query("name", "null")

	// Get IMAGE_FOLDER_PATH var in .env file
	folderPath := os.Getenv("IMAGE_FOLDER_PATH")

	// Build the full path (./folder/imageName)
	imageFilePath := filepath.Join(folderPath, name)

	// If the image doesn't exist, fetch it and store it
	if _, err := os.Stat(imageFilePath); errors.Is(err, os.ErrNotExist) {
		// Get IMAGE_PROVIDER var in .env file (https://thispersondoesnotexist.com)
		imageProvider := os.Getenv("IMAGE_PROVIDER")

		// Get the image
		resp, err := http.Get(imageProvider)

		if err != nil {
			log.Fatalln("Error loading the image provider :", err)
		}

		// Create the destination file
		fs, err := os.Create(imageFilePath)

		if err != nil {
			log.Fatalln("Error creating file destination :", err)
		}

		// Copy the http response into the destination file
		_, err = io.Copy(fs, resp.Body)
		if err != nil {
			log.Fatalln("Error copy into file destination :", err)
		}
	}

	return c.SendFile(imageFilePath)
}
