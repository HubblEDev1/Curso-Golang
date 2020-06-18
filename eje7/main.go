package main

//Funciones y sobrecar
import "fmt"

func main() {
	//numero, estado := uno(1)
	//fmt.Printf("%d %v", numero, estado)
	fmt.Println(dos(1))
	fmt.Println(calculo(3, 8, 2))
}

func uno(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {()
		return 10, true
	}
}

func dos(numero int) (z int) {
	if numero == 1 {
		z = numero
	}
	return
}

func calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		total = total + num
	}
	return total
}
