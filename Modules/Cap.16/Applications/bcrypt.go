package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	x := "teste123"
	senhaErrada := "teste1234"
	sb, err := bcrypt.GenerateFromPassword([]byte(x), 10)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaErrada)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(x)))
}
