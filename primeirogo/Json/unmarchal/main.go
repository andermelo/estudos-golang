package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome": "Rex", "raca": "Poodle", "idade": 4}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorroEmJson2 := `{"nome": "Pipoca", "raca": "Pitbull", "idade": 10}`

	c2 := make(map[string]any)

	if erro := json.Unmarshal([]byte(cachorroEmJson2), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
