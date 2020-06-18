package main

import (
	"fmt"
)

func main() {
	//scaner := bufio.NewScanner(os.Stdin)
	//scaner.Scan()
	//texto := scaner.Text()
	//lista := strings.Split(texto, " ")
	fmt.Println(Sum(4, 1, 6))
}

func Sum(num ...int) (z int) {
	total := 0
	for _, numero := range num {
		total = total + numero
	}
	z = total
	return
}
