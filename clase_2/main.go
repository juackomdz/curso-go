package main

//"flag"
//"errors"
//"fmt"
//"log"
//"os"
//"math/rand"
//"time"

import (
	modulo "clase_2/modulo_ejemplo"
	"fmt"
	"log"
)

func main() {

	fmt.Println(modulo.Saludar())

	fmt.Println(modulo.Ejempl2("joaquin"))

	log.Fatalln("error de ejecucion")
	////logs
	/*
		//err := errors.New("error de prueba")
		//log.Fatal(err)
		//fmt.Println("texto de prueba")

		//log.Println("Starting app......")
		log.Fatalln("error de ejecucion")
		//log.Panicln("error de ejecucion")


		f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		log.SetOutput(f)
		log.Println("error de ejecucion")

	*/
	/*
		//modulo os / argumentos

		nombre := flag.String("nombre", "", "el nombre")
		edad := flag.Int("edad", 19, "la edad")
		flag.Parse()
		fmt.Println("tu nombre es:", *nombre, "y tu edad es:", *edad)

	*/
	/*
		//math  random
		aleatorio := rand.Intn(100)
		fmt.Println(aleatorio)

		min := 5
		max := 10

		//seed esta deprecado
		rand.Seed(time.Now().UnixNano())
		aleatorio2 := rand.Intn(max-min) + min
		fmt.Println(aleatorio2)

		rand.NewSource(time.Now().UnixNano())
		aleatorio3 := rand.Intn(max-min) + min
		fmt.Println(aleatorio3)
	*/
	///time
	/*
		fmt.Println(time.Now())
		fecha := time.Now()
		fmt.Println("el año es", fecha.Year())
		fmt.Println("el mes es", int(fecha.Month()))
		fmt.Println("el dia es", fecha.Day())
		//fecha formateada
		fmt.Println(fecha.Day(), "/", int(fecha.Month()), "/", fecha.Year())
	*/

	///strings
	/*
		cadena := "mi muñeca me hablo"
		fmt.Println(cadena)
		fmt.Println(strings.ToUpper(cadena))

		letras := strings.Split(cadena, "")
		fmt.Println(letras)

		pos := strings.Index(cadena, "muñeca")
		fmt.Println(pos)
		if pos == -1 {
			fmt.Println("no se encontro la palabra")
		} else {
			fmt.Println("la palabra se encontro en la posicion", pos)
		}

		repeat := strings.Repeat(cadena, 4)
		fmt.Println(repeat)

		nueva := strings.Replace(cadena, "mi", "la", -1)
		fmt.Println(nueva)

		fmt.Println(string(nueva[0:6]))
	*/
}
