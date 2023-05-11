package main

import "fmt"

func main() {
	fmt.Println("\r\n\r\nContando de trás pra frente")
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
