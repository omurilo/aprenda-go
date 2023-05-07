package main

import "fmt"

func main() {
	persons := map[string][]string{
		"murilo_alves":   {"Programar", "Jogar"},
		"gabriela_souza": {"Malhar", "Vender"},
		"luiza_pessoa":   {"Brincar", "Brincar mais"},
	}

	persons["eleni_fernandes"] = []string{"Ler", "Cozinhar", "Dormir"}

	for key, value := range persons {
		fmt.Println(key)
		for i, hobby := range value {
			fmt.Println("\t", i, " - ", hobby)
		}
	}
}
