//Interfaces
package main

import "fmt"

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	Carnivoro() bool
}

type vegetal interface {
	clasificacion() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

type mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

func (h *hombre) respirar()    { h.respirando = true }
func (h *hombre) comer()       { h.comiendo = true }
func (h *hombre) pensar()      { h.pensando = true }
func (h *hombre) sexo() string { return "hombre" }

func (m *mujer) respirar()    { m.respirando = true }
func (m *mujer) comer()       { m.comiendo = true }
func (m *mujer) pensar()      { m.pensando = true }
func (m *mujer) sexo() string { return "mujer" }

func HumansResp(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y respiro \n", hu.sexo())
}

func main() {
	pedro := new(hombre)
	HumansResp(pedro)
	maria := new(mujer)
	HumansResp(maria)
}
