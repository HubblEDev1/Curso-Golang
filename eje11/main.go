//vectores y arreglos
package main

import "fmt"

//var tabla [10]bool
//*****************************************************************************************************
//vectores dinamicos

func main() {
	//tabla := []int{2, 1, 4, 3, 5, 2, 9, 5, 3, 7}
	//var tabla [4][7]int
	//matriz := []int{2, 5, 2}
	//fmt.Println(matriz)
	//vector()
	vector1()
	vector2()
}

func vector() {
	elementos := [5]int{4, 2, 8, 11, 7}
	porcion := elementos[1:3]
	fmt.Println(porcion)
}

func vector1() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func vector2() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i*i)
	}
	fmt.Println(nums)
}
