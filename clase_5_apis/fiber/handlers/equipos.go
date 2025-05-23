package handlers

import (
	"clase_5_fiber/db"
	"clase_5_fiber/dtos"
	"clase_5_fiber/modelos"
	"context"

	"github.com/go-rel/rel"
	"github.com/gofiber/fiber/v2"
)

func Get_equipos(ctx *fiber.Ctx) error {

	var datos []modelos.EquipoModel

	if err := db.Connect().FindAll(context.TODO(), &datos); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Error al obtener los datos"})
	}
	return ctx.Status(200).JSON(datos)
}

func Get_equipo_id(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var dato modelos.EquipoModel

	if err := db.Connect().Find(context.TODO(), &dato, rel.Eq("id", id)); err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Error al obtener los datos"})
	}
	return ctx.Status(200).JSON(dato)
}

func Post_equipos(ctx *fiber.Ctx) error {

	var dto dtos.EquiposDTO
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Error al parsear los datos"})
	}

	dato := modelos.EquipoModel{
		Nombre: dto.Nombre,
		Liga:   dto.Liga,
	}
	if err := db.Connect().Insert(context.TODO(), &dato); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Error al insertar los datos"})
	}

	return ctx.Status(201).JSON(fiber.Map{"mensaje": "creado con exito"})
}

func Put_equipo(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var dto dtos.EquiposDTO
	var dato modelos.EquipoModel

	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Error al parsear los datos"})
	}

	if err := db.Connect().Find(context.TODO(), &dato, rel.Eq("id", id)); err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Error al obtener los datos"})
	}

	dato.Nombre = dto.Nombre
	dato.Liga = dto.Liga

	if err := db.Connect().Update(context.TODO(), &dato); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Error al actualizar los datos"})
	}

	return ctx.Status(200).JSON(fiber.Map{"mensaje": "actualizado con exito"})
}

func Delete_equipo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var dato modelos.EquipoModel

	if err := db.Connect().Find(context.TODO(), &dato, rel.Eq("id", id)); err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Error al obtener los datos"})
	}

	if err := db.Connect().Delete(context.TODO(), &dato); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Error al eliminar los datos"})
	}

	return ctx.Status(200).JSON(fiber.Map{"mensaje": "eliminado con exito"})
}
