package main

import "fmt"

func main() {
	x := 387

	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(x)
}
