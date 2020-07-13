package main

//Funciones anonimas, los closures son funciones creadas para proteger código
import (
	"fmt"
)

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	//fmt.Println(Suma(1, 3, 4, 7))
	fmt.Printf("%d", Calculo(5, 7))
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("%d", Calculo(5, 7))

	Operaciones()
	/*Closures*/
	tabla := 2
	miTabla := Tabla(tabla)
	for i := 1; i <= 10; i++ {
		fmt.Println(miTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		return 5
	}
	fmt.Println(resultado())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia
	}
}

/*
func Suma(numero ...int) int {
	total := 0
	for _, entero := range numero {
		total += entero
	}
	return total
}
*/
