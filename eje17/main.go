package main

import (
	"log"
)

func main() {
	//archivo := "prueba.txt"
	//f, err := os.Open(archivo)
	//defer f.Close() //Cierra archivo defer ejecuta esa parte de codigo antes de dalir

	ejePanic()

}

//Manejo de report y panic
func ejePanic() {
	defer func() { //defer con funcion anonima
		r := recover()
		if r != nil {
			log.Fatalf("Ocurrio un erro %v", r)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
