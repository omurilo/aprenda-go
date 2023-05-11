package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	client1 := cliente{
		nome:      "Fulano",
		sobrenome: "Da Silva",
		fumante:   false,
	}

	client2 := cliente{"Murilo", "Alves", false}

	fmt.Println(client1, client2)
}
