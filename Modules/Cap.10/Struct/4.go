package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Murilo",
		idade: 28,
	}

	fmt.Println(x)
}
