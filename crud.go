package main

import (
	"fmt"
)

func main() {
	/*e := Estudiante{
		Active: true,
	}

	err := Crear(e)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creado exitosamente!")
	*/

	es, err := Consultar()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(es)
}
