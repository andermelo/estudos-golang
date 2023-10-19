package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Pedro"
	u.idade = 4
	fmt.Println(u)

	usuario2 := usuario{nome: "Pedro"}
	fmt.Println(usuario2)
}
