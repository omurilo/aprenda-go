package main

import "fmt"

func main() {
	x := returnfuncao(2)

	y := x(3)

	fmt.Println(y)
	fmt.Println(returnfuncao(1)(15))
}

func returnfuncao(x int) func(int) int {
	return func(i int) int {
		return x * i * 10
	}
}
