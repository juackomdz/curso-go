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

	err := database.CategoriaColle.FindOne(ctx, filtro).Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(404, echo.Map{"mensaje": "elemento no encontrado"})
	} else if err != nil {
		log.Printf("%v", err)
	}

	return c.JSON(200, result)
}

func Categoria_put(c echo.Context) error {

	var cate dto.CategoriaDTO

	objId, _ := bson.ObjectIDFromHex(c.Param("id"))

	if err := c.Bind(&cate); err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid input"})
	}

	if len(cate.Nombre) == 0 {
		return c.JSON(400, echo.Map{"error": "Ingrese un nombre"})
	}

	//para modificar mas datos usar un map

	registro := map[string]interface{}{
		"nombre": cate.Nombre,
	}
	update := bson.D{{"$set", registro}}
	var result bson.D
	//registro := bson.D{{"$set", bson.D{{Key: "nombre", Value: cate.Nombre}}}}
	if err := database.CategoriaColle.FindOneAndUpdate(context.TODO(), bson.D{{"_id", objId}}, update).Decode(&result); err != nil {
		log.Printf("%v", err.Error())
	}
	return c.JSON(200, echo.Map{"mensaje": "actualizado con exito"})
}

func Categoria_delete(c echo.Context) error {

	objId, _ := bson.ObjectIDFromHex(c.Param("id"))

	filtro := bson.D{{"_id", objId}}

	_, err := database.CategoriaColle.DeleteOne(context.TODO(), filtro)
	if err != nil {
		return c.JSON(400, echo.Map{"mensaje": "algo salio mal"})
	}
	return c.JSON(200, echo.Map{"mensaje": "se elimino correctamente"})
}
