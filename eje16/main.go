package main

import "fmt"

func main() {
	n := 2
	recursion(n)

}

func recursion(numero int) {
	if numero > 100 {
		return
	}
	fmt.Println(numero)
	recursion(numero * 2)

}
