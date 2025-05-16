package handlers

import (
	"clase_5/gin/database"
	db "clase_5/gin/database"
	"clase_5/gin/modelos"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Get_pelicula(c *gin.Context) {

	db := bun.NewDB(database.Conexion(), pgdialect.New())
	var datos []modelos.PeliculaModel

	err := db.NewSelect().Model(&datos).Relation("Tematica").Scan(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}
	c.JSON(http.StatusOK, datos)
}

func Get_pelicula_id(c *gin.Context) {

	db := bun.NewDB(db.Conexion(), pgdialect.New())

	id := c.Param("id")

	exist, erre := db.NewSelect().Model((*modelos.PeliculaModel)(nil)).Where("id=?", id).Exists(context.TODO())
	if erre != nil {
		log.Fatal(erre.Error())
	}

	if !exist {
		c.JSON(400, gin.H{"mensaje": "error inesperado"})
		return
	}
	var dato modelos.PeliculaModel

	err := db.NewSelect().Model(&dato).Relation("Tematica").Where("p.id=?", id).Scan(context.TODO())

	if err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}
	c.JSON(200, dato)
}

func Post_pelicula(c *gin.Context) {

	c.JSON(201, gin.H{"mensaje": "creado con exito"})
}
