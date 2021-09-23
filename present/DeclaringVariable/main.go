package main

import (
	"fmt"
)

//KEYWORD IDENTIFIER TYPE
var z int

func main() {
	x := 42 + 7
	y := "James Bond"
	z = 5
	fmt.Println(x)
	fmt.Println(y)
	x = 50

	i, err := fmt.Println(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T %d\n", z, z)
	fmt.Printf("%T\n%d\n", x, i)
}
