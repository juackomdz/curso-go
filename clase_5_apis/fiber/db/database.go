package db

import (
	"database/sql"
	"log"

	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	_ "github.com/jackc/pgx/v5/stdlib"
	//_ "github.com/lib/pq"
)

func Connect() rel.Repository {

	//dr := sql.Drivers()

	//log.Println("driver=", dr)

	//-----usar libreria sql standar porque por defecto no realiza el cambio de driver si se usa pgx--------------
	dsn := "postgres://postgres:pass123@localhost:5432/rel?sslmode=disable"
	conn, err := sql.Open("pgx", dsn)

	if err != nil {
		//panic(err)
		log.Println("err: " + err.Error())
	}

	adapter := postgres.New(conn)

	repo := rel.New(adapter)
	return repo
}
