package handlers

import (
	"clase_3/db"
	"clase_3/modelos"
	"context"
	"fmt"
	"log"
	"os"
)

func Listar() {

	db.Conectar()
	sql := "select * from clientes;"

	//metodo nativo databse/sql
	rows, err := db.Db.Query(context.Background(), sql)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	defer db.CerraConexion()
	///mostrar datos sin formato
	/*
		datos := modelos.Clientes{}
		for rows.Next() {
			dato := modelos.Cliente{}
			rows.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
			//fmt.Println(&dato.Nombre)
			datos = append(datos, dato)
		}
		fmt.Println(datos)
		fmt.Println("--------------------")
	*/

	///mostrar datos con formato

	for rows.Next() {
		var dato = modelos.Cliente{}
		err := rows.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %v | nombre: %s | correo: %s | telefono: %s\n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
	}

	//metodo propio de driver postgre
	//data, _ := pgx.CollectRows(rows, pgx.RowToStructByName[modelos.Cliente])
	/*
		fmt.Println("----------------------------------------------------")
		for _, cs := range data {
			fmt.Printf("id: %v | nombre: %s | correo: %s | telefono: %s\n", cs.Id, cs.Nombre, cs.Correo, cs.Telefono)
		}
	*/
}
