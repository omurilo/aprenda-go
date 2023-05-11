package main

import "fmt"

func main() {
	basica()

	total, quantos := soma(10, 10, 1, 3, 5, 7)
	fmt.Println(total, quantos)
}

func basica() {
	fmt.Println("Oi, bom dia!")
}

func soma(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma, len(x)
}
