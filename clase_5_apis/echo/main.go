package main

import (
	"clase_5_echo/database"
	"clase_5_echo/handlers"
	cjwt "clase_5_echo/jwt"

	_ "clase_5_echo/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			swagger ejemplo
// @description	This is a sample server celler server.
// @version		1.0
//
// @host			localhost:8085
// @BasePath		/
func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

	m := e.Group("/mongo")
	m.POST("/categoria", handlers.Categoria_post)
	m.GET("/categoria", handlers.Categoria_get)
	m.GET("/categoria/:id", handlers.Categoria_get_id)
	m.PUT("/categoria/:id", handlers.Categoria_put)
	m.DELETE("/categoria/:id", handlers.Categoria_delete)

	m.POST("/producto", handlers.Producto_post)
	m.GET("/producto", handlers.Producto_get)
	m.GET("/producto/:id", handlers.Producto_get_id)
	m.PUT("/producto/:id", handlers.Producto_put)
	m.DELETE("/producto/:id", handlers.Producto_delete)

	//------------------------jwt----------------------
	m.POST("/usuario", handlers.Registro_user)
	m.POST("/login", handlers.Login)

	r := e.Group("/secure", cjwt.MiddlewareJWT())

	r.GET("/usuarios", handlers.Listar_user_secure)

	//---------------------api externa-----------------------
	e.GET("/externo", handlers.Traer_datos)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8085"))
}
