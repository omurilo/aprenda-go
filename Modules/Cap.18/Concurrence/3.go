package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	contador := 0
	totalDeGoRoutines := 1000

	wg.Add(totalDeGoRoutines)

	var mu sync.Mutex

	for i := 0; i < totalDeGoRoutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}
