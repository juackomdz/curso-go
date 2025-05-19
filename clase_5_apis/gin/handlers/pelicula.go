package handlers

import (
	"clase_5/gin/database"
	db "clase_5/gin/database"
	"clase_5/gin/dto"
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

	db := bun.NewDB(db.Conexion(), pgdialect.New())
	var body dto.PeliculaDTO
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(400, gin.H{"mensaje": "error en el body"})
		return
	}

	if len(body.Nombre) == 0 {
		c.JSON(400, gin.H{"mensaje": "debe ingresar el nombre"})
		return
	}

	_, err := db.NewInsert().Model(&body).Exec(context.TODO())

	if err != nil {
		c.JSON(400, gin.H{"mensaje": "hubo un error inesperado"})
		return
	}
	c.JSON(201, gin.H{"mensaje": "creado con exito"})
}

func Put_pelicula(c *gin.Context) {

	db := bun.NewDB(db.Conexion(), pgdialect.New())
	var body dto.PeliculaDTO

	id := c.Param("id")

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"mensaje": "error en el body"})
		return
	}

	model := modelos.PeliculaModel{Nombre: body.Nombre, Descripcion: body.Descripcion, TematicaId: body.TematicaId}
	_, err := db.NewUpdate().Model(&model).Where("id=?", id).Exec(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"mensaje": "actualizado con exito",
	})
}

func Delete_pelicula(c *gin.Context) {

	db := bun.NewDB(db.Conexion(), pgdialect.New())
	var pelicula modelos.PeliculaModel
	id := c.Param("id")

	_, err := db.NewDelete().Model(&pelicula).Where("id=?", id).Exec(context.TODO())

	if err != nil {
		c.JSON(400, gin.H{"mensaje": "hubo un error inesperado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "borraro con exito"})
}
