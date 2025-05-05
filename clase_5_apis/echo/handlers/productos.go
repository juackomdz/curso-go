package handlers

import (
	"clase_5_echo/database"
	"clase_5_echo/dto"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Producto_post(c echo.Context) error {

	var produ dto.ProductoDTO

	if err := c.Bind(&produ); err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid input"})
	}

	if len(produ.Nombre) == 0 {
		return c.JSON(400, echo.Map{"error": "Ingrese un nombre"})
	}
	cateId, _ := bson.ObjectIDFromHex(produ.CategoriaId)
	registro := bson.D{{"nombre", produ.Nombre},
		{"precio", produ.Precio},
		{"descripcion", produ.Descripcion},
		{"stock", produ.Stock},
		{"categoria_id", cateId},
	}

	database.ProductoColle.InsertOne(context.TODO(), registro)

	return c.JSON(http.StatusCreated, echo.Map{
		"mensaje": "Creado con exito",
	})
}

func Producto_get(c echo.Context) error {

	pipeline := []bson.M{
		bson.M{"$match": bson.M{}},
		bson.M{"$lookup": bson.M{
			"from":         "categorias",
			"localField":   "categoria_id",
			"foreignField": "_id",
			"as":           "categoria",
		},
		},
	}

	cursor, err := database.ProductoColle.Aggregate(context.TODO(), pipeline)
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

func Producto_get_id(c echo.Context) error {

	objId, _ := bson.ObjectIDFromHex(c.Param("id"))

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"_id": objId}},
		bson.M{"$lookup": bson.M{
			"from":         "categorias",
			"localField":   "categoria_id",
			"foreignField": "_id",
			"as":           "categoria",
		},
		},
	}

	cursor, err := database.ProductoColle.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	if result == nil {
		return c.JSON(400, echo.Map{"error": "no hay datos para mostrar"})
	}

	return c.JSON(200, result[0])
}

func Producto_put(c echo.Context) error {

	var produ dto.ProductoDTO

	objId, _ := bson.ObjectIDFromHex(c.Param("id"))

	if err := c.Bind(&produ); err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid input"})
	}

	if len(produ.Nombre) == 0 {
		return c.JSON(400, echo.Map{"error": "Ingrese un nombre"})
	}

	cateId, _ := bson.ObjectIDFromHex(produ.CategoriaId)
	//para modificar mas datos usar un map
	registro := map[string]interface{}{
		"nombre":       produ.Nombre,
		"precio":       produ.Precio,
		"descripcion":  produ.Descripcion,
		"stock":        produ.Stock,
		"categoria_id": cateId,
	}
	update := bson.D{{"$set", registro}}
	var result bson.D

	if err := database.ProductoColle.FindOneAndUpdate(context.TODO(), bson.D{{"_id", objId}}, update).Decode(&result); err != nil {
		log.Printf("%v", err.Error())
	}

	return c.JSON(200, echo.Map{"mensaje": "actualizado con exito"})
}

func Producto_delete(c echo.Context) error {

	objId, _ := bson.ObjectIDFromHex(c.Param("id"))

	filtro := bson.D{{"_id", objId}}

	res, err := database.ProductoColle.DeleteOne(context.TODO(), filtro)
	if err != nil {
		return c.JSON(400, echo.Map{"mensaje": "algo salio mal"})
	}

	if res.DeletedCount == 0 {
		return c.JSON(404, echo.Map{"mensaje": "elemento no encontrado"})
	}
	return c.JSON(200, echo.Map{"mensaje": "se elimino correctamente"})
}
