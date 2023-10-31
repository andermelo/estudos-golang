package main

import (
	"api/src/router/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API RODANDO -------")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8080", r))
}
