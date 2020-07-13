package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leer()
	leer2()
	escritura()
	escritura1()
}
func leer() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(string(archivo))
	}

}

func leer2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")

		}
	}
	archivo.Close()
}

func escritura() {
	archivo, err := os.Create("./new.txt")
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Fprintln(archivo, "cadena de texto")
	archivo.Close()
}

func escritura1() {
	fileName := "./ne.txt"
	if append(fileName, "\nEsta es una segunda linea") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	_, err = arch.WriteString(texto) //_ Indica que no nos interesa almacenar en una variable un dato
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}
