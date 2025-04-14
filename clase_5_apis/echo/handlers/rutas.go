package handlers

import "github.com/labstack/echo/v4"

func Ejemplo_get(c echo.Context) error {

	return c.JSON(200, "Echo con get")
}

func Ejemplo_post(c echo.Context) error {

	return c.JSON(200, "Echo con post")
}

func Ejemplo_put(c echo.Context) error {

	return c.JSON(200, "Echo con put")
}

func Ejemplo_delete(c echo.Context) error {

	return c.JSON(200, "Echo con delete")
}

func Ejemplo_get_params(c echo.Context) error {

	id := c.Param("id")
	return c.JSON(200, "Echo con get | id: "+id)
}

func Ejemplo_get_query_string(c echo.Context) error {

	id := c.QueryParam("id")
	return c.JSON(200, "Echo con get | id: "+id)
}
