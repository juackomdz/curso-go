package handlers

import (
	"clase_5_fiber/db"
	"clase_5_fiber/dtos"
	"clase_5_fiber/modelos"
	"context"

	"github.com/go-rel/rel"
	"github.com/gofiber/fiber/v2"
)

func Get_jugadores(ctx *fiber.Ctx) error {

	var datos []modelos.JugadorModel

	if err := db.Connect().FindAll(context.TODO(), &datos); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"mensaje": "Error al obtener los datos"})
	}

	if err := db.Connect().Preload(context.TODO(), &datos, "equipo"); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"mensaje": "Error al relacionar"})
	}

	return ctx.Status(200).JSON(datos)
}

func Post_jugadores(ctx *fiber.Ctx) error {

	var dto dtos.JugadoresDTO
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Error al parsear los datos"})
	}

	dato := modelos.JugadorModel{
		Nombre:   dto.Nombre,
		Posicion: dto.Posicion,
		EquipoId: dto.EquipoId,
	}
	if err := db.Connect().Insert(context.TODO(), &dato); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Error al insertar los datos"})
	}

	return ctx.Status(201).JSON(fiber.Map{"mensaje": "creado con exito"})
}

func Get_jugador_id(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var dato modelos.JugadorModel

	if err := db.Connect().Find(context.TODO(), &dato, rel.Eq("id", id)); err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Error al obtener los datos"})
	}

	if err := db.Connect().Preload(context.TODO(), &dato, "equipo"); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"mensaje": "Error al relacionar"})
	}

	return ctx.Status(200).JSON(dato)
}

func Put_jugador(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var dto dtos.JugadoresDTO
	var dato modelos.JugadorModel

	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Error al parsear los datos"})
	}

	if err := db.Connect().Find(context.TODO(), &dato, rel.Eq("id", id)); err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Error al obtener los datos"})
	}

	dato.Nombre = dto.Nombre
	dato.Posicion = dto.Posicion
	dato.EquipoId = dto.EquipoId

	if err := db.Connect().Update(context.TODO(), &dato); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Error al actualizar los datos"})
	}

	return ctx.Status(200).JSON(fiber.Map{"mensaje": "actualizado con exito"})
}

func Delete_jugador(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var dato modelos.JugadorModel
	if errDato := db.Connect().Find(context.TODO(), &dato, rel.Eq("id", id)); errDato != nil {
		return ctx.Status(404).JSON(fiber.Map{"mensaje": "no se encontro el registro"})
	}

	if err := db.Connect().Delete(context.TODO(), &dato); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"mensaje": "no se pudo eliminar"})
	}

	return ctx.Status(200).JSON(fiber.Map{"mensaje": "eliminado con exito"})
}
