package main

import (
	"fmt"
)

func main() {
	err := Borrar(2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Borrado correctamente")
}
