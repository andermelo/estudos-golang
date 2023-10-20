package main

import (
	"linha-de-comando/appl"
	"log"
	"os"
)

func main() {

	aplicacao := appl.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
