package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Esta es una cadena concatenada")
	my_var := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	my_var.Scan()
	fmt.Println(my_var.Text() + scanner.Text())

}
