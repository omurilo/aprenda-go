package main

import "fmt"

type person struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	person
	dentesarrancados int
	salário          float64
}

type arquiteto struct {
	person
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei", x.dentesarrancados, "dentes e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
}

func main() {
	mrdente := dentista{
		person: person{
			nome:      "Mister Dente",
			sobrenome: "Da Silva",
			idade:     37,
		},
		dentesarrancados: 8973,
		salário:          98878887898,
	}

	mrtijolo := arquiteto{
		person: person{
			nome:      "Mister Tijolo",
			sobrenome: "Da Silva",
			idade:     52,
		},
		tipodeconstrução: "Hospícios",
		tamanhodaloucura: "Demais",
	}

	// mrdente.oibomdia()
	// mrtijolo.oibomdia()

	serhumano(mrdente)
	serhumano(mrtijolo)
}
