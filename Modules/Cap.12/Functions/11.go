package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scanln(&input)

	for input != 0 {
		x := f91(input)
		fmt.Printf("f91(%v) = %v\n", input, x)
		fmt.Scanln(&input)
	}
}

func f91(x int) int {
	if x <= 100 {
		return f91(f91(x + 11))
	}

	return x - 10
}
