package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		bola string = "Bola"
		nome string = "Pedro"
	)

	fmt.Println(nome, "gosta de jogar", bola)

	idade, altura := 4, 1.21

	fmt.Println(nome, "tem", idade, "anos", "e", altura, "de altura")

}
