package main

import "fmt"

func main() {

	/*
		fmt.Println("Hola mundo con GO")

		//variables
		var nombre string = "Joaquin"
		fmt.Println(nombre)

		nombre2 := "Juan"
		fmt.Println(nombre2)

		//constantes
		//-- %s funciona con string, %v funciona con int
		const constante = 457
		fmt.Printf("constante contiene: %v \n", constante)

		//tipos de datos
		var string1 string = "texto"
		fmt.Println(string1)

		textoGrande := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
		fmt.Println(textoGrande)

		var estado bool = true
		fmt.Println(estado)

		//concatenacion
		var edad int = 28
		fmt.Println("la edad es:", edad)

		var flotante float32 = 23.45
		fmt.Println(flotante)

		//refelct y typeof
		string2 := 3434343.4546
		fmt.Println(reflect.TypeOf(string2))

		//punteros
		//me indica el espacio de memoria
		var puntero string = "texto"
		fmt.Println(&puntero)
		//inicializa el valor como nulo == nil
		/*

			var numero *int
			fmt.Println(numero)


		//condicionales
		edad2 := 14
		if edad2 >= 18 {
			fmt.Println("eres mayor de edad")
		} else {
			fmt.Println("eres menor de edad")
		}

		color := "verde"
		if color == "rojo" {
			fmt.Println("el color es rojo")
		} else if color == "azul" {
			fmt.Println("el color es azul")
		} else {
			fmt.Println("el color no es rojo ni azul")
		}
		//operador logico
		if color == "verde" && edad2 != 14 {
			fmt.Println("el color es verde y la edad no es 14")
		} else {
			fmt.Println("revise el color y la edad")
		}
		//switch case
		switch color {
		case "rojo":
			fmt.Println("el color es rojo")
		case "azul":
			fmt.Println("el color es azul")
		default:
			fmt.Println("el color no es rojo ni azul")
		}

		//declarar variable en condicion
		if variable := 2; variable == 1 {
			fmt.Println("la variable es 1")
		} else {
			fmt.Println("la variable no es 1")
		}

		//ciclos e iteraciones
		i := 1
		for i < 11 {
			//i++   cuenta hasta 11
			fmt.Println(i)
			i++ //  cuenta hasta 10
		}
		fmt.Println("------------------------------------")
		for i2 := 1; i2 < 11; i2++ {
			fmt.Println(i2)
		}

		//slices
		//slices son arrays dinamicos

		//arreglos

		var paises [4]string
		paises[0] = "chile"
		paises[1] = "argentina"
		paises[2] = "uruguay"
		paises[3] = "paraguay"

		fmt.Println(paises)
		fmt.Println(paises[0])
		fmt.Println("el largo es", len(paises))

		fmt.Println("------------------------------")
		//slices
		var paises2 = make([]string, 5)
		paises2[0] = "chile"
		paises2[1] = "alemania"
		paises2[2] = "españa"
		paises2[3] = "italia"
		paises2[4] = "bolivia"
		fmt.Println(paises2)

		//agregar elemento a slice
		paises2 = append(paises2, "bulgaria")
		fmt.Println(paises2)
		//eliminar elemento a slice

		paises2 = append(paises2[:5], paises2[5+1:]...)
		fmt.Println(paises2)

		//maps
		paises3 := make(map[string]string)
		paises3["chile"] = "santiago"
		paises3["argentina"] = "buenos aires"
		paises3["uruguay"] = "montevideo"
		fmt.Println(paises3)
		fmt.Println(paises3["uruguay"])

		fmt.Println("-----------------------------")

		paises4 := map[int]string{
			1: "chile",
			2: "argentina",
			3: "uruguay",
		}
		fmt.Println(paises4)
		//validacion si existe en el map
		pais, existe := paises4[12]
		if existe {
			fmt.Println("el pais existe", pais)
		} else {
			fmt.Println("el pais no existe")
		}
		//eliminar elemento
		delete(paises4, 1)
		fmt.Println(paises4)

		//recorrer map con ciclo for
		for i, v := range paises4 {
			fmt.Println(i, v)
			if v == "uruguay" {
				fmt.Println("clave=", i)
			}
		}
	*/

	///funciones

	/*
		miFuncion()
		miFuncConParemetros(2, 5)
		fmt.Println(miFuncConReturn("Joaquin"))
		nombre, apelido, edad := miFuncMultiple("Juan", "perez", 30)
		fmt.Println("Hola mi nombre es=", nombre, apelido, "y tengo:", edad)
		fmt.Println("la multiplicacion es:", multi(2, 5))
		tabla := tabla(2)

		for i := 1; i < 10; i++ {
			fmt.Println("2 *", i, "=", tabla())
		}
	*/

	/*
		//ejemplo 1
		fmt.Println(miFuncion("juan"))
		time.Sleep(time.Second * 5)
		fmt.Println(miFuncion("pablo"))

		//ejemplo 2
		canal := make(chan string)
		go func() {
			canal <- miFuncion("joaquin")
		}()

		fmt.Println(<-canal)
	*/

	///recursividad
	//funcion(2)

	//errores
	//mifun()

	/*
		//estructuras
		//forma sin utlizar puntero --- create bd
		estr := Persona{
			Id:     1,
			Nombre: "joaquin",
			Edad:   27,
			Correo: "joaquin@corre.cl",
		}

		fmt.Println(estr)

		//forma utilizando puntero --- update bd
		p := new(Persona)
		p.Id = 2
		p.Nombre = "juan"
		p.Edad = 30
		p.Correo = "juan@correo.cl"

		fmt.Println(reflect.TypeOf(p))
		fmt.Println(p.Nombre)

		fmt.Println("------------------------------")

		cate := Categoria{Id: 1, Nombre: "categoria 1", Slug: "slug-cate-1"}
		produ := Producto{Id: 1, Nombre: "produ-1", Precio: 12345, CategoriaId: cate}

		fmt.Printf("%+v", produ)
	*/

	e := Estructura{}
	fmt.Println(e.miFunc())
	fmt.Println(e.Calculo(3, 5))

	e.campo1 = "texto"
	fmt.Println(e.miFunc())

}

///interfaces
/*

 */

type Estructura struct {
	campo1 string
}

func (e *Estructura) miFunc() string {
	return "hola mundo con interface" + e.campo1
}

func (*Estructura) Calculo(n1 int, n2 int) int {
	return n1 + n2
}

/*
// //estructuras
type Persona struct {
	Id     int
	Nombre string
	Edad   int
	Correo string
}

// /estructura anidada
type Categoria struct {
	Id     int
	Nombre string
	Slug   string
}

type Producto struct {
	Id          int
	Nombre      string
	Precio      int
	CategoriaId Categoria
}

*/
// //manejo de errores
// defer y panic
/*
func mifun() {
	defer fmt.Println("mensaje final")
	fmt.Println("mensaje 1")
	a := 1
	if a == 1 {
		panic("errofmt")
	}
}
*/
/*
func funcion(valor int) {
	dato := valor + 1
	fmt.Println(dato)
	if dato < 10 {
		funcion(dato)
	}
}
*/
// gorutinas
/*
func miFuncion(nombre string) string {
	return "Hola " + nombre
}
*/
/*
func miFuncion() {
	fmt.Println("Hola desde funcion")
}

func miFuncConParemetros(n1 int, n2 int) {
	fmt.Println("el total de la suma es =", n1+n2)
}

func miFuncConReturn(nombre string) string {
	return "mi nombre es = " + nombre
}

func miFuncMultiple(nombre string, apellido string, edad int) (string, string, int) {
	return nombre, apellido, edad
}

//funciones anonimas
var multi = func(n1 int, n2 int) int {
	return n1 * n2
}

//funciones clousure
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
*/
