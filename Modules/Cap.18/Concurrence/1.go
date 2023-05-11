package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// add(total functions)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Add(2)
	go println("func(0):")
	go println("func(1):")
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func println(x string) {
	for i := 0; i <= 100; i++ {
		fmt.Println(x, i)
		time.Sleep(20)
	}
	wg.Done()
}
