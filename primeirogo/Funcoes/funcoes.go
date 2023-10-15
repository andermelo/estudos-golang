package main

import "fmt"

func calcular(n1 int8, n2 int8) (int8, int8) {
	return n1 + n2, n1 - n2

}

func main() {
	resultSoma, resultSubtracao := calcular(55, 45)

	fmt.Println(resultSoma, resultSubtracao)

	var f = func(txt string) {
		println("minha funcao F:", txt)
	}

	f("bola")
}
