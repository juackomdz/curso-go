package main

import (
	"clase_5_echo/database"
	"clase_5_echo/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	database.BoolConectado()

	e.GET("/ejemplo", handlers.Ejemplo_get)
	e.POST("/ejemplo", handlers.Ejemplo_post)
	e.PUT("/ejemplo", handlers.Ejemplo_put)
	e.DELETE("/ejemplo", handlers.Ejemplo_delete)

	e.GET("/ejemplo/:id", handlers.Ejemplo_get_params)
	e.GET("/query-string", handlers.Ejemplo_get_query_string)

	e.GET("/users", handlers.Get_users)
	e.POST("/users", handlers.Post_users)
	e.GET("/users/:id", handlers.Get_id)
	e.PUT("/users/:id", handlers.Update_user)
	e.DELETE("/users/:id", handlers.Delete_user)

	//---------------mongodb----------------------
	prefix := "/mongo"
	e.POST(prefix+"/categoria", handlers.Categoria_post)

	e.Logger.Fatal(e.Start(":8085"))
}
