package main

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(
		fiber.Config{
			Prefork:       true,
			CaseSensitive: true,
			BodyLimit:     20 * 1024 * 1024,
			StrictRouting: true,
			ServerHeader:  "Fiber",
			AppName:       "Test App v1.0.1",
		})
	app.Use(logger.New())
	app.Use(cors.New())

	app.Post("/", func(c *fiber.Ctx) error {
		fileheader, err := c.FormFile("picture")
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)
		}

		file, err := fileheader.Open()
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		defer file.Close()

		buffer, err := io.ReadAll(file)
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)
		}

		errDir := createFolder("uploads")
		if errDir != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}

		filename, err := imageProcessing(buffer, 40, "uploads")
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}
		return c.JSON(fiber.Map{
			"picture": "/uploads/" + filename,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		fileheader, err := c.FormFile("picture")
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}

		file, err := fileheader.Open()
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}
		defer file.Close()

		buffer, err := io.ReadAll(file)
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}

		errDir := createFolder("uploads")
		if errDir != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}

		filename, err := imageProcessing(buffer, 40, "uploads")
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}
		return c.JSON(fiber.Map{
			"picture": "/uploads/" + filename,
		})
	})

	app.Listen(":3939")
}
