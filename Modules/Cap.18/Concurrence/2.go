package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	contador := 0
	totalDeGoRoutines := 50

	wg.Add(totalDeGoRoutines)

	for i := 0; i < totalDeGoRoutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}
