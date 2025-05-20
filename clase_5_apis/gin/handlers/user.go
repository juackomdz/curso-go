package handlers

import (
	"clase_5/gin/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Password string `json:"password"`
}

type UserDTO struct {
	Nombre   string `json:"nombre"`
	Password string `json:"password"`
}

var users []User
var uid int

func Usuarios(c *gin.Context) {

	if len(users) == 0 {
		c.JSON(404, gin.H{"mensaje": "No hay usuarios"})
		return
	}
	c.JSON(200, users)
}

func Login(c *gin.Context) {

	var u UserDTO

	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, err.Error())
		return
	}

	passByte := []byte(u.Password)
	passBD := []byte(users[0].Password)

	if err := bcrypt.CompareHashAndPassword(passBD, passByte); err != nil {
		c.JSON(400, gin.H{"mensaje": "credenciales incorrectas"})
		return
	}

	jwtKey, err := jwt.Generar_token(u.Nombre, strconv.Itoa(uid))
	if err != nil {
		c.JSON(400, gin.H{"mensaje": "error generando el token"})
		return
	}

	c.JSON(200, gin.H{"user": users[0].Nombre, "token": jwtKey})
}

func Register(c *gin.Context) {

	var u UserDTO
	uid++

	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, err.Error())
		return
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	us := User{
		Id:       uid,
		Nombre:   u.Nombre,
		Password: string(bytes),
	}
	users = append(users, us)

	c.JSON(200, gin.H{"mensaje": "creado con exito"})
}
