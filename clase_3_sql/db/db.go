package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var Db *pgx.Conn

// metodo para conectar
func Conectar() {
	errors := godotenv.Load()
	if errors != nil {
		log.Fatal(errors)
	}

	//connStr := "postgres://username:password@localhost:5432/database_name"
	connStr1 := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME")
	db, err := pgx.Connect(context.Background(), connStr1)
	if err != nil {
		log.Fatal(err)
	}

	Db = db

}

func CerraConexion() {
	Db.Close(context.Background())
}
