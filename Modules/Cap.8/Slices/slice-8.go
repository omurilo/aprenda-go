package main

import "fmt"

func main() {
	primeiroslice := []int{1, 2, 3, 4, 5}

	fmt.Println(primeiroslice)

	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)
	// a pegadinha de fuder os dev, o primeiro slice n é imutável e o append vai
	// substituir os valores do segundo slice no primeiro slice.
	fmt.Println(primeiroslice)
}
