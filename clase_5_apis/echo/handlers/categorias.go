package handlers

import (
	"clase_5_echo/database"
	"clase_5_echo/dto"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Categoria_post(c echo.Context) error {

	var cate dto.CategoriaDTO

	if err := c.Bind(&cate); err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid input"})
	}

	registro := bson.D{{Key: "nombre", Value: cate.Nombre}}

	database.CategoriaColle.InsertOne(context.TODO(), registro)

	return c.JSON(http.StatusCreated, echo.Map{
		"mensaje": "Creado con exito",
	})
}
