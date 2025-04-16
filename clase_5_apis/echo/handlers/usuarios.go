package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var uid int

type User struct {
	Id       int       `json:"id"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Edad     int       `json:"edad"`
	Fecha    time.Time `json:"fecha"`
}

type UserDTO struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var users []User

func Get_users(c echo.Context) error {
	if len(users) == 0 {
		return c.JSON(404, echo.Map{"mensaje": "No hay usuarios"})
	}
	return c.JSON(200, users)
}

func Post_users(c echo.Context) error {
	var u UserDTO

	uid++
	if err := c.Bind(&u); err != nil {
		return c.JSON(400, err.Error())
	}

	us := User{
		Id:       uid,
		Nombre:   u.Nombre,
		Apellido: u.Apellido,
		Edad:     u.Edad,
		Fecha:    time.Now()}
	users = append(users, us)

	return c.JSON(200, echo.Map{"mensaje": "creado con exito"})
}

func Get_id(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, v := range users {
		if v.Id == id {
			return c.JSON(200, v)
		}
	}
	return c.JSON(404, echo.Map{"mensaje": "usuario no encontrado"})
}

func Update_user(c echo.Context) error {

	var u UserDTO
	id, _ := strconv.Atoi(c.Param("id"))

	for i, _ := range users {
		if err := c.Bind(&u); err != nil {
			return c.JSON(400, echo.Map{
				"mensaje": err.Error(),
			})
		}

		if users[i].Id == id {
			users[i].Nombre = u.Nombre
			users[i].Apellido = u.Apellido
			users[i].Edad = u.Edad

			return c.JSON(200, echo.Map{"mensaje": "Actualizado con exito"})
		} else {
			return c.JSON(404, echo.Map{"mensaje": "usuario no encontrado"})
		}
	}

	return c.JSON(http.StatusBadRequest, echo.Map{"mensaje": "error"})
}

func Delete_user(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, v := range users {
		if v.Id == id {
			users = append(users[:i], users[i+1:]...)

			return c.JSON(200, echo.Map{
				"mensaje": "usuario eliminado con exito",
			})
		}
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
		"mensaje": "error",
	})
}
