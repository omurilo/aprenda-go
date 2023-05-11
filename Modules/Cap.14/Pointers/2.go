package main

import "fmt"

func main() {
	x := 11

	receiveValue(x)
	// receivePointer(&x)

	fmt.Println("Na 'main':", x)
}

func receiveValue(x int) {
	x++
	fmt.Println("Na função:", x)
}

func receivePointer(x *int) {
	*x++
	fmt.Println("Na função pointer:", *x)
}
