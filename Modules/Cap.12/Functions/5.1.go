package main

import "fmt"

func serhumano2(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho", g.(dentista).salário)
	case arquiteto:
		fmt.Println("Eu construo", g.(arquiteto).tipodeconstrução)
	}
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

	serhumano2(mrtijolo)
	fmt.Println("")
	serhumano2(mrdente)
}
