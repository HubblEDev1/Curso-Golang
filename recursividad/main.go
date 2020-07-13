package main

import "fmt"

func main() {
	var n int = 0
	potencia(n)
}

func potencia(num int) {
	if num > 100 {
		return
	} else {
		fmt.Println(num)
		num += 2
	}
	potencia(num)
}
