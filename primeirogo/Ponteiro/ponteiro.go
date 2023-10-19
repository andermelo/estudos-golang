package main

import "fmt"

func main() {
	var variavel int
	var ponteiro *int

	variavel = 100
	ponteiro = &variavel

	fmt.Println(variavel, ponteiro)
	fmt.Println("----------------")
	variavel += 14
	fmt.Println(variavel, *ponteiro)

}
