package main

import (
	"encoding/json"
	"fmt"
)

type pessoa2 struct {
	Nome          string  `json:"nome"`
	Sobrenome     string  `json:"sobrenome"`
	Idade         int     `json:"idade"`
	Profissao     string  `json:"profissao"`
	Contabancaria float64 `json:"contabancaria"`
}

func main() {
	sb := []byte(`{"nome":"James","sobrenome":"Bond","idade":40,"profissao":"Agente Secreto","contabancaria":1000000.5}`)

	var jamesbond pessoa2
	err := json.Unmarshal(sb, &jamesbond)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(jamesbond.Profissao)
}
