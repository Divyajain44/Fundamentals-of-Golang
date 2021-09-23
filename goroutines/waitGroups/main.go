package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// ch := make(chan int)

	wg.Add(2)
	go fun1()
	go fun2()

	wg.Wait()
	fmt.Println("exiting main")

}

func fun1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n", i)
	}
	wg.Done()
}

func fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("in fun2")
	}
	wg.Done()
}
