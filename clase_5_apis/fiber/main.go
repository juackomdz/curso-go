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
	v1.Get("/equipos/:id", handlers.Get_equipo_id)
	v1.Post("/equipos", handlers.Post_equipos)
	v1.Put("/equipos/:id", handlers.Put_equipo)
	v1.Delete("/equipos/:id", handlers.Delete_equipo)

	v1.Get("/jugadores", handlers.Get_jugadores)
	v1.Post("/jugadores", handlers.Post_jugadores)
	v1.Put("/jugadores/:id", handlers.Put_jugador)
	v1.Delete("jugadores/:id", handlers.Delete_jugador)
	v1.Get("/jugadores/:id", handlers.Get_jugador_id)

	app.Listen(":8087")
}
