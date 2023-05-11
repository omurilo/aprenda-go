package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

// Método = função anexada à um tipo
func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "dia bom dia!")
}

func main() {
	mauricio := pessoa{"Mauricio", 30}

	mauricio.oibomdia()
}
