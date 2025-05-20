package main

import (
	"clase_5/gin/handlers"
	"clase_5/gin/jwt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

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
	v1.POST("/pelicula", handlers.Post_pelicula)
	v1.PUT("/pelicula/:id", handlers.Put_pelicula)
	v1.DELETE("/pelicula/:id", handlers.Delete_pelicula)

	//------------------------jwt----------------------

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
	v2 := r.Group("/api/v2")
	v2.Use(jwt.MiddlewareJWT())
	{
		v2.GET("/users", handlers.Usuarios)
	}

	r.Run(":8086")
}
