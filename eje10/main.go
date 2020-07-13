package main

import (
	"fmt"
)

func main() {
	n := 4
	mi_Tabla := Tablas(n)

	for i := 1; i <= 10; i++ {
		fmt.Println(mi_Tabla())
	}
}

func Tablas(num int) func() int {
	numero := num
	sec := 0
	return func() int {
		sec++
		return numero * sec
	}
}
