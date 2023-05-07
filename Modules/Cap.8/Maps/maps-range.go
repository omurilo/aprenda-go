package main

import "fmt"

func main() {
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)

	amigos["gopher"] = 444444

	for key, value := range amigos {
		fmt.Println(key, value)
	}

	delete(amigos, "gopher")
	fmt.Println("Gopher is deleted", amigos)
}
