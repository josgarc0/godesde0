package main

import (
	//defer_panic "github.com/josgarc0/godesde0/defer"
	//"github.com/josgarc0/godesde0/modelos"

	"github.com/josgarc0/godesde0/middleware"
	//"github.com/josgarc0/godesde0/iteraciones"
)

func main() {
	/*//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)


	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	estado, texto := ejercicios.RestoVariables("dd")
	fmt.Println(estado)
	fmt.Println(texto)
	*/
	//teclado.IngresoNumeros()

	//iteraciones.Iterar()

	//fmt.Println(ejercicios.TabladeMultiplicar())

	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo2()
	////funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(2)

	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.Capacidad()

	///mapas.MostrarMapas()
	//users.AltaUsuario()
	//Pedro := new(modelos.Hombre)
	//ejer_interfaces.HumanoRespirando(Pedro)

	//Maria := new(modelos.Mujer)
	//ejer_interfaces.HumanoRespirando(Maria)

	//defer_panic.EjemploPanic()
	/*canal1 := make(chan bool)
	go goroutines.MiNombreLento("jose manuel", canal1)
	fmt.Println("Estoy aqui")

	<-canal1 */

	//webserver.MiWebServer()

	middleware.MiMiddleware()

}
