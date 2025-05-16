package handlers

import (
	db "clase_5/gin/database"
	"clase_5/gin/dto"
	"clase_5/gin/modelos"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Get_tematica(c *gin.Context) {

	db := bun.NewDB(db.Conexion(), pgdialect.New())

	var datos []modelos.TematicaModel
	err := db.NewSelect().Model(&datos).Scan(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{"mensaje": err})
		return
	}
	c.JSON(http.StatusOK, datos)
}

func Get_tematica_id(c *gin.Context) {

	db := bun.NewDB(db.Conexion(), pgdialect.New())

	id := c.Param("id")
	var dato modelos.TematicaModel

	err := db.NewSelect().Model(&dato).Where("id=?", id).Scan(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}
	c.JSON(200, dato)
}

func Post_tematica(c *gin.Context) {

	db := bun.NewDB(db.Conexion(), pgdialect.New())

	var dto dto.TematicaDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}

	_, err := db.NewInsert().Model(&dto).Exec(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "creado con exito"})
}

func Put_tematica(c *gin.Context) {

	db := bun.NewDB(db.Conexion(), pgdialect.New())
	id := c.Param("id")

	var dto dto.TematicaDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}

	_, err := db.NewUpdate().Model(&dto).Where("id=?", id).Exec(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "actualizado con exito"})
}

func Delete_tematica(c *gin.Context) {

	db := bun.NewDB(db.Conexion(), pgdialect.New())
	id := c.Param("id")

	var dato modelos.TematicaModel
	_, err := db.NewDelete().Model(&dato).Where("id=?", id).Exec(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{"mensaje": err.Error()})
		return
	}
	c.JSON(200, gin.H{"mensaje": "eliminado con exito"})
}
