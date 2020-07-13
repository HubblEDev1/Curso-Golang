package user

import "time"

type Usuario struct {
	Id     int
	Name   string
	Fecha  time.Time
	Status bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fecha time.Time, status bool) {
	this.Id = id
	this.Name = nombre
	this.Fecha = fecha
	this.Status = status
}
