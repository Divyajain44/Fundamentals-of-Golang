package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
	sliceLength := len(slice)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	fmt.Println("Running for loop…")
	for i := 0; i < sliceLength; i++ {

		go func(i int) {
			fmt.Println(runtime.NumGoroutine())
			defer wg.Done()
			val := slice[i]
			fmt.Printf("i: %v, val: %v\n", i, val)

		}(i)
	}

	wg.Wait()
	fmt.Println("Finished for loop")
}
