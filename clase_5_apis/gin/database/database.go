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

	//-----------borra y crea la tabla pero no crea la foreign key-------------
	/*
		db := bun.NewDB(sql, pgdialect.New())

		if err := db.ResetModel(context.TODO(), &modelos.TematicaModel{}); err != nil {
			log.Fatal(err)
		}
	*/
	//-----------borrar la tabla y crearla nuevamente para agregar foreign key------------
	/*
		//db := bun.NewDB(sql, pgdialect.New())
		_, errd := db.NewDropTable().Model((*modelos.PeliculaModel)(nil)).IfExists().Exec(context.TODO())
		if errd != nil {
			log.Fatal(errd)
		}
		_, errt := db.NewCreateTable().Model((*modelos.PeliculaModel)(nil)).WithForeignKeys().Exec(context.TODO())
		if errt != nil {
			log.Fatal(errt)
		}
	*/
	return sql
}
