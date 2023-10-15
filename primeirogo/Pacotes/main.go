package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("oi Golang pelo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("anderson@gmail.com")

	fmt.Println(erro)
}
