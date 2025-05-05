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
	e.GET(prefix+"/categoria", handlers.Categoria_get)
	e.GET(prefix+"/categoria/:id", handlers.Categoria_get_id)
	e.PUT(prefix+"/categoria/:id", handlers.Categoria_put)
	e.DELETE(prefix+"/categoria/:id", handlers.Categoria_delete)

	e.POST(prefix+"/producto", handlers.Producto_post)
	e.GET(prefix+"/producto", handlers.Producto_get)
	e.GET(prefix+"/producto/:id", handlers.Producto_get_id)
	e.PUT(prefix+"/producto/:id", handlers.Producto_put)
	e.DELETE(prefix+"/producto/:id", handlers.Producto_delete)

	e.Logger.Fatal(e.Start(":8085"))
}
