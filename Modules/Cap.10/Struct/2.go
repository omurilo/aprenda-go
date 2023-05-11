package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Murilo",
			idade: 28,
		},
		titulo:  "Programador",
		salario: 20100,
	}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.nome)

}
