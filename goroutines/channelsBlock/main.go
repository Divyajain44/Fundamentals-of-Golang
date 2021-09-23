package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("in func")
		<-ch

	}()
	fmt.Println("exiting main")

}
