package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	jsonString := `{"nome":"James","sobrenome":"Bond","idade":40,"profissao":"Agente Secreto","contabancaria":1000000.5}`

	// decoder não precisa de um []bytes, pode receber uma string ou uma stream de strings
	decoder := json.NewDecoder(strings.NewReader((jsonString)))
	var jamesbond pessoa
	decoder.Decode(&jamesbond)
	fmt.Println(jamesbond)
}
