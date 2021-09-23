package main

import (
	"fmt"
)

func main() {
	for {
		go anotherGoroutine()
		// forever()
	}
}

// func forever() {
// 	for {
// 	}
// }
func anotherGoroutine() {
	i := 0
	for {
		i++
		fmt.Print(i)
		fmt.Print("\n")
	}
}
