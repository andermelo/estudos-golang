package main

import "fmt"

type user struct {
	nome  string
	idade uint8
}

func (u user) save() {
	fmt.Printf("usuario %s salvo\n", u.nome)
}

func (u *user) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := user{nome: "Pedro", idade: 4}
	fmt.Println(usuario1)
	usuario1.save()

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
