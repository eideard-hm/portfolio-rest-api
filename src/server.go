package src

import (
	"github.com/eideard-hm/portfolio-rest-api/src/controllers"
	"github.com/eideard-hm/portfolio-rest-api/src/database"
	"github.com/gofiber/fiber/v2"
)

func Server() *fiber.App {
	// start database connection
	database.DbConnection()
	// exec migrations
	database.Migrate()

	app := fiber.New()

	app.Get("/", controllers.GetUsersHandler)

	return app
}
