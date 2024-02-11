package routers

import (
	"github.com/NikolaJankovikj/go-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func GenerateApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	gameGroup := app.Group("/games")
	gameGroup.Get("/", handlers.GetGames)
	gameGroup.Post("/", handlers.CreateGame)
	gameGroup.Delete("/:id", handlers.DeleteGame)
	gameGroup.Put("/:id", handlers.UpdateGame)

	return app
}
