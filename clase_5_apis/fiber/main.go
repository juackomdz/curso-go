package main

import (
	"clase_5_fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/hola", func(c *fiber.Ctx) error {
		return c.Status(400).JSON(fiber.Map{"mensaje": "Hola mundo"})
	})

	//---------------grupo de rutas--------------
	v1 := app.Group("/api/v1")
	v1.Get("/equipos", handlers.Get_equipos)

	app.Listen(":8087")
}
