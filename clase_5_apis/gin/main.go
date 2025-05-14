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

	r.Run(":8086")
}
