package main

import (
	"clase_5/gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"mensaje": "pong"})
	})

	v1 := r.Group("/api/v1")
	v1.GET("/tematica", handlers.Get_tematica)
	v1.GET("/tematica/:id", handlers.Get_tematica_id)
	v1.POST("/tematica", handlers.Post_tematica)
	v1.PUT("/tematica/:id", handlers.Put_tematica)
	v1.DELETE("/tematica/:id", handlers.Delete_tematica)

	v1.GET("/pelicula", handlers.Get_pelicula)
	v1.GET("/pelicula/:id", handlers.Get_pelicula_id)

	r.Run(":8086")
}
