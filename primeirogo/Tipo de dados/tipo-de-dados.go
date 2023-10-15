package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 1000000000000000
	fmt.Println(numero)

	var erro error = errors.New("Erro interno!")

	fmt.Println(erro)
}
