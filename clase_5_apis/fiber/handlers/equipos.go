package handlers

import (
	"clase_5_fiber/db"
	"clase_5_fiber/modelos"
	"context"

	"github.com/gofiber/fiber/v2"
)

func Get_equipos(ctx *fiber.Ctx) error {

	var datos []modelos.EquipoModel

	if err := db.Connect().FindAll(context.TODO(), &datos); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Error al obtener los datos"})
	}
	return ctx.Status(200).JSON(datos)
}
