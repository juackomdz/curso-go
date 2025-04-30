package handlers

import (
	"clase_5_echo/database"
	"clase_5_echo/dto"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Categoria_post(c echo.Context) error {

	var cate dto.CategoriaDTO

	if err := c.Bind(&cate); err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid input"})
	}

	if len(cate.Nombre) == 0 {
		return c.JSON(400, echo.Map{"error": "Ingrese un nombre"})
	}
	registro := bson.D{{Key: "nombre", Value: cate.Nombre}}

	database.CategoriaColle.InsertOne(context.TODO(), registro)

	return c.JSON(http.StatusCreated, echo.Map{
		"mensaje": "Creado con exito",
	})
}

func Categoria_get(c echo.Context) error {

	cursor, err := database.CategoriaColle.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var data []bson.M
	if err = cursor.All(context.TODO(), &data); err != nil {
		log.Fatal(err)
	}

	if data == nil {
		return c.JSON(400, echo.Map{"error": "no hay datos para mostrar"})
	}
	return c.JSON(200, data)
}

func Categoria_get_id(c echo.Context) error {

	objId, _ := bson.ObjectIDFromHex(c.Param("id"))

	filtro := bson.D{{"_id", objId}}

	var result bson.M

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	log.Printf("%v", objId)

	err := database.CategoriaColle.FindOne(ctx, filtro).Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(404, echo.Map{"mensaje": err.Error()})
	} else if err != nil {
		log.Printf("%v", err)
	}
	//log.Printf()
	/*
		if err := database.CategoriaColle.FindOne(context.TODO(), filtro).Decode(&result); err != nil {
			return c.JSON(404, echo.Map{"mensaje": err.Error()})
		}
	*/
	return c.JSON(200, result)
}
