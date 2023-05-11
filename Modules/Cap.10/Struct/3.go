package main

import "fmt"

type professional struct {
	person  pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := professional{
		person: pessoa{
			nome:  "Murilo",
			idade: 28,
		},
		titulo:  "Programador",
		salario: 20100,
	}

	pessoa3 := professional{pessoa{"Maricota", 31}, "Pizzaiola", 10000}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.person.nome)
	fmt.Println(pessoa3.person.nome)

}
