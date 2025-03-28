package handlers

import (
	"bufio"
	"clase_3/db"
	"clase_3/modelos"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Listar() {

	db.Conectar()
	sql := "select * from clientes order by id asc"

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

func ListarPorId(id int) {
	db.Conectar()

	sql := "select * from clientes where id = $1"
	rows, err := db.Db.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatalln(err)
	}

	data, _ := pgx.CollectRows(rows, pgx.RowToStructByName[modelos.Cliente])

	fmt.Println("--------------------------------------------------")
	for _, data := range data {
		fmt.Printf("Id: %v | Nombre: %s | Correo: %s | Telefono: %s", data.Id, data.Nombre, data.Correo, data.Telefono)
	}

	defer db.CerraConexion()
}

func Guardar(cli modelos.Cliente) {
	db.Conectar()

	sql := "insert into clientes(nombre, correo, telefono) values ($1, $2, $3)"
	res, err := db.Db.Exec(context.Background(), sql, cli.Nombre, cli.Correo, cli.Telefono)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.CerraConexion()

	fmt.Println(res)
	fmt.Println("Creado con exito")
}

func Editar(id int, cli modelos.Cliente) {
	db.Conectar()

	sql := "update clientes set nombre=$1, correo=$2, telefono=$3 where id=$4"

	res, err := db.Db.Exec(context.Background(), sql, cli.Nombre, cli.Correo, cli.Telefono, id)
	if err != nil {
		log.Fatal(err)
	}

	defer db.CerraConexion()

	fmt.Println(res)
	fmt.Println("Editado con exito")
}

func Eliminar(id int) {
	db.Conectar()

	sql := "delete from clientes where id=$1"

	res, err := db.Db.Exec(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	defer db.CerraConexion()

	fmt.Println(res)
	fmt.Println("Eliminado con exito")
}

// /////////////////FUNCIONES///////////////
var id int

var nombre, correo, telefono string

func Ejecutar() {
	scan := bufio.NewScanner(os.Stdin)

	fmt.Println("Seleccione una opcion: ")
	fmt.Println("1. Listar")
	fmt.Println("2. Buscar por id")
	fmt.Println("3. Crear")
	fmt.Println("4. Editar")
	fmt.Println("5. Eliminar")

	if scan.Scan() {
		for {

			if scan.Text() == "1" {
				Listar()
				return
			}
			if scan.Text() == "2" {
				fmt.Println("Ingrese el id a buscar")
				//forma de leer 1

				fmt.Scanln(&id)
				ListarPorId(id)
				return

				//forma de leer 2
				/*
					if scan.Scan() {
						id, _ = strconv.Atoi(scan.Text())
					}
					ListarPorId(id)
					return
				*/
			}
			if scan.Text() == "3" {
				fmt.Print("Ingrese nombre: ")
				fmt.Scanln(&nombre)

				fmt.Print("Ingrese el correo: ")
				fmt.Scanln(&correo)

				fmt.Print("Ingrese telefono: ")
				fmt.Scanln(&telefono)

				cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Guardar(cliente)
				return

			}

			if scan.Text() == "4" {
				fmt.Print("Ingrese el id a editar: ")
				fmt.Scanln(&id)

				fmt.Print("Ingrese el nuevo nombre: ")
				fmt.Scanln(&nombre)

				fmt.Print("Ingrese el nuevo correo: ")
				fmt.Scanln(&correo)

				fmt.Print("Ingrese el nuevo telefono: ")
				fmt.Scanln(&telefono)

				cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Editar(id, cliente)
				return
			}

			if scan.Text() == "5" {
				fmt.Print("Ingrese el id a eliminar: ")
				fmt.Scanln(&id)

				Eliminar(id)
				return
			}

		}
	}
}
