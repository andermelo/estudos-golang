package main

import "fmt"

func main() {
	canal := make(chan string, 3)
	canal <- "Olá mundo!"
	canal <- "Olá mundo 2!"
	canal <- "Olá mundo 3!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
