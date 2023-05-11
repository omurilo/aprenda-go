package main

import "fmt"

func main() {
	si := []int{10, 10, 1, 2, 3, 5}
	total := soma2(si...)
	fmt.Println(total)

	total2 := soma2()
	fmt.Println(total2)
}

func soma2(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma
}
