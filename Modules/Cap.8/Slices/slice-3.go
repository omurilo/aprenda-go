package main

import "fmt"

func main() {
	// 											0						1							2							3
	sabores := []string{"peperoni", "mozzarella", "margarita", "quatro queijos"}
	fatia := sabores[2:len(sabores)]
	fmt.Println(fatia)

	for i := 0; i < len(sabores); i++ {
		fmt.Println("o item na posiçao", i, "tem o valor", sabores[i])
	}

	fatia2 := sabores[:]
	fmt.Println("fatia 2 é:", fatia2)
}
