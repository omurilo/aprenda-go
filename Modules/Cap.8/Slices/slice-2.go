package main

import "fmt"

func main() {
	slice := []string{"banana", "maça", "jaca", "pêssego"}
	slice = append(slice, "melancia")
	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}

	slice[5] = "uva"
	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}
}
