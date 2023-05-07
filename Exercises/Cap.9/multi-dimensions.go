package main

import "fmt"

func main() {
	persons := [][]string{
		{
			"Murilo",
			"Alves",
			"Programar",
		},
		{
			"Gabriela",
			"Souza",
			"Vender cosméticos",
		},

		{
			"Luiza",
			"Pessoa",
			"Brincar",
		},
	}

	fmt.Println(persons, "\n\n")

	for _, value := range persons {
		fmt.Println(value[0])
		for _, item := range value {
			fmt.Println("\t", item)
		}
	}
}
