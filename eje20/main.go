package main

import (
	"fmt"
)

var result int

/*Son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben y
devuelven los mismos tipos de de variable*/
func main() {
	fmt.Println("Inicio")

	result = opsMid(sumar)(2, 6)
	fmt.Println(result)
	result = opsMid(restar)(7, 4)
	fmt.Println(result)
	result = opsMid(multiplicar)(8, 4)
	fmt.Println(result)

}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func opsMid(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
