package pkg2

import (
	"fmt"
)

func Pkg2(x ...string) {
	for _, v := range x {
		fmt.Println(v)
	}
}
