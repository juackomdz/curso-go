package handlers

import (
	db "clase_5/gin/database"
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
	}
	c.JSON(http.StatusOK, datos)
}
