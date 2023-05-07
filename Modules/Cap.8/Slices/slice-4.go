package main

import "fmt"

func main() {
	sabores := []string{"peperoni", "mozzarella", "margarita", "quatro queijos"}
	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)
}
