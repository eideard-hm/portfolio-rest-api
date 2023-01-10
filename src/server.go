package src

import (
	"github.com/eideard-hm/portfolio-rest-api/src/database"
	"github.com/gofiber/fiber/v2"
)

func Server() *fiber.App {
	// start database connection
	database.DbConnection()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "Hello Word",
		})
	})

	return app
}
