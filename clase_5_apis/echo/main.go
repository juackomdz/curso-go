package main

import (
	"clase_5_echo/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/ejemplo", handlers.Ejemplo_get)
	e.POST("/ejemplo", handlers.Ejemplo_post)
	e.PUT("/ejemplo", handlers.Ejemplo_put)
	e.DELETE("/ejemplo", handlers.Ejemplo_delete)

	e.GET("/ejemplo/:id", handlers.Ejemplo_get_params)
	e.GET("/query-string", handlers.Ejemplo_get_query_string)

	e.Logger.Fatal(e.Start(":8085"))
}
