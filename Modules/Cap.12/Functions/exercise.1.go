package main

import "fmt"

func main() {
	t := someImpares(soma2, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println("Soma somente os ímpares:", t)
}

func someImpares(f func(x ...int) int, y ...int) int {
	var slice []int

	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}

	return f(slice...)
}
