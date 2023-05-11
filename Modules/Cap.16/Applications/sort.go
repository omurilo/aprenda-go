package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"uva", "banana", "manga", "laranja", "pêssego"}
	si := []int{51, 50, 1, 2, 3, 4, 6, 7, 9, 5, 8}
	sort.Strings(ss)
	sort.Ints(si)

	fmt.Println(ss)
	fmt.Println(si)
}
