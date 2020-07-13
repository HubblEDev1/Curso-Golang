package main

import "fmt"

func main() {
	pais := make(map[string]string, 5)
	fmt.Println(pais)
	pais["Mex"] = "D.F"
	fmt.Println(pais["Mex"])

	champ := map[string]int{
		"Guatemala": 31,
		"Barcelona": 11,
		"UNAM":      85}
	champ["Bayern"] = 34
	delete(champ, "UNAM")
	fmt.Println(champ)

	for equipo, puntaje := range champ {
		fmt.Println(equipo, puntaje)
	}
	pun, existe := champ["Barcelona"]
	fmt.Println(pun, existe)
}
