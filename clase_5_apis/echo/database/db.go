package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var cliente = Conectar()
var db = "mongo_echo"

var CategoriaColle = cliente.Database(db).Collection("categorias")
var ProductoColle = cliente.Database(db).Collection("productos")
var UsersColle = cliente.Database(db).Collection("usuarios")

func Conectar() *mongo.Client {

	// Conectar a la base de datos

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017/" + db))

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("CONEXION EXITOSA....")
	return client
}

func BoolConectado() int {
	err := cliente.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
