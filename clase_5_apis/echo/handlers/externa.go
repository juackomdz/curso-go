package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"io"

	"github.com/labstack/echo/v4"
)

type BodyResponse struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

type pokemon struct {
	id              int
	name            string
	base_experience int
	is_default      bool
}

func Traer_datos(c echo.Context) error {

	var response BodyResponse
	//var data pokemon
	url := "https://jsonplaceholder.typicode.com/todos/4"
	//url := "https://pokeapi.co/api/v2/pokemon/1"

	res, err := http.Get(url)

	if err != nil {
		return c.JSON(400, echo.Map{"mensaje": "error"})
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(400, echo.Map{"mensaje": "error al formatear"})
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("error unmarshal")
	}
	log.Println(res.Status)
	return c.JSON(200, echo.Map{
		"id":        response.Id,
		"userid":    response.UserId,
		"title":     response.Title,
		"completed": response.Completed,
	})
}
