package main

import "fmt"

func generico(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generico("string")
	generico(1.5)
	generico(299232)
	generico(true)

	// exemplo de aplicaco do tipo generico para interface e p fmt.Println("string", 132, true)
	fmt.Println("string", 132, true)
}
