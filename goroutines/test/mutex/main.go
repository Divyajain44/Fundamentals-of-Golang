package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	var x sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			x.Lock()
			v := counter
			// time.Sleep(time.Second)
			// runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			x.Unlock()
			wg.Done()
		}()
		// fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
