package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	jamesbond := pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000000.50,
	}

	darthvader := pessoa{"Arthur", "Skywalker", 50, "Vilão Intergalático", 500000000.50}

	j, err := json.Marshal(jamesbond)

	if err != nil {
		fmt.Println("Error:", err)
	}

	d, err := json.Marshal(darthvader)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%v\n\n", string(j))
	fmt.Printf("%v", string(d))
}
