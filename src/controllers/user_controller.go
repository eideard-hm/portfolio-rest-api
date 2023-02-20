package controllers

import "github.com/gofiber/fiber/v2"

func GetUsersHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "Retrieve all users",
	})
}
