package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	var contador int64 = 0
	totalDeGoRoutines := 50

	wg.Add(totalDeGoRoutines)

	for i := 0; i < totalDeGoRoutines; i++ {
		go func() {
			runtime.Gosched()
			atomic.AddInt64(&contador, 1)
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}
