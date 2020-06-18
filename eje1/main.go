//Variables
package main

import (
	"fmt"
	"strconv"
)

/*
type point struct {
	x, y int
}
*/
func main() {
	//p := point{1, 2}
	//fmt.Printf("%v\n", p)
	n1, n2, n3, text := 2, 5, 3, "Hola a todos"
	var moneda float64 = 0
	text = strconv.Itoa(int(moneda))
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(text)

}
