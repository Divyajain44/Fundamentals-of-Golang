package main

import "fmt"

func main() {
	fmt.Println("hello world")
	n, err := fmt.Println("Hello, playground", 4+2, true)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(2 + 2)
}
