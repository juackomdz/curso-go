package handlers

import (
	"clase_5_echo/database"
	"clase_5_echo/dto"
	jwt "clase_5_echo/jwt"
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func Registro_user(c echo.Context) error {
	var body dto.UserDTO

	if err := c.Bind(&body); err != nil {
		return c.JSON(400, echo.Map{"mensaje": "un error inesperado"})
	}

	if len(body.Username) == 0 && len(body.Password) == 0 {
		return c.JSON(400, echo.Map{"mensaje": "debe ingresar un username y un password"})
	}

	var user bson.M

	if err := database.UsersColle.FindOne(context.TODO(), bson.M{"correo": body.Correo}).Decode(&user); err == nil {
		return c.JSON(400, echo.Map{"mensaje": "correo ya existe"})
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	registro := bson.D{{"username", body.Username}, {"password", string(bytes)}, {"correo", body.Correo}}

	database.UsersColle.InsertOne(context.TODO(), registro)

	return c.JSON(201, echo.Map{"mensaje": "creado con exito"})
}

func Login(c echo.Context) error {

	var req dto.LoginDTO

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, echo.Map{"mensaje": "un error inesperado"})
	}

	if len(req.Username) == 0 {
		return c.JSON(400, echo.Map{"mensaje": "debe ingresar un username"})
	}
	if len(req.Password) == 0 {
		return c.JSON(400, echo.Map{"mensaje": "debe ingresar el password"})
	}

	var user bson.M

	if err := database.UsersColle.FindOne(context.TODO(), bson.M{"username": req.Username}).Decode(&user); err != nil {
		return c.JSON(400, echo.Map{"mensaje": "usuario no encontrado"})
	}

	passByte := []byte(req.Password)
	passBD := []byte(user["password"].(string))

	if errPasswd := bcrypt.CompareHashAndPassword(passBD, passByte); errPasswd != nil {
		return c.JSON(400, echo.Map{"mensaje": "credenciales incorrectas"})
	} else {
		strObj := user["_id"].(bson.ObjectID).Hex()
		jwtKey, err := jwt.Generar_token(user["correo"].(string), user["username"].(string), strObj)
		if err != nil {
			return c.JSON(400, echo.Map{"mensaje": "error generando token"})
		}
		return c.JSON(200, echo.Map{"user": user["username"].(string), "token": jwtKey})
	}

}

func Listar_user_secure(c echo.Context) error {

	//claims := c.Get("claims").(*Claims)
	cursor, err := database.UsersColle.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var data []bson.M
	if err = cursor.All(context.TODO(), &data); err != nil {
		log.Fatal(err)
	}

	if data == nil {
		return c.JSON(400, echo.Map{"error": "no hay datos para mostrar"})
	}
	return c.JSON(200, data)
}
