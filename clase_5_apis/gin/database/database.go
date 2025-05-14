package database

import (
	"database/sql"
	"log"

	_ "github.com/uptrace/bun/driver/pgdriver"
)

func Conexion() *sql.DB {

	dsn := "postgres://postgres:pass123@localhost:5432/bun?sslmode=disable"
	sql, err := sql.Open("pg", dsn)
	if err != nil {
		log.Fatal(err)
	}
	/*
		db := bun.NewDB(sql, pgdialect.New())
		if err := db.ResetModel(context.TODO(), &modelos.TematicaModel{}); err != nil {
			log.Fatal(err)
		}
	*/
	return sql
}
