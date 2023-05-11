package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int           { return Len(x) }
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (x ordenarPorPotencia) Swap(i, j int)      { Swap(x, i, j) }

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int           { return Len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
func (x ordenarPorConsumo) Swap(i, j int)      { Swap(x, i, j) }

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int           { return Len(x) }
func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (x ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { Swap(x, i, j) }

func Len(x []carro) int        { return len(x) }
func Swap(x []carro, i, j int) { x[i], x[j] = x[j], x[i] }

func main() {
	carros := []carro{
		{
			nome:     "Fusca",
			potencia: 30, // 30cv
			consumo:  8,  // 8 km/L
		},
		{"BMW X3", 180, 13},
		{"Mercedes A140", 82, 18},
		{"Chevette", 50, 8},
		{"Porsche", 300, 3},
	}

	fmt.Println(carros)

	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println("Ordenados por potência 0 < 1\n", carros, "\n")
	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println("Ordenados por consumo 0 < 1\n", carros, "\n")
	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println("Ordenados por lucro pro dono do posto 0 > 1\n", carros, "\n")
}
