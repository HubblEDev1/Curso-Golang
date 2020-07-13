package main

import (
	"fmt"
	"time"

	us "./user"
)

type ed struct {
	us.Usuario
}

func main() {
	user := new(ed)
	user.AltaUsuario(14, "Edwin", time.Now(), true)
	fmt.Println(user.Usuario)
}
