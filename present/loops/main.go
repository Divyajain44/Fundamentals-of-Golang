package main

import (
	"fmt"
)

func main() {
	x := 1
	for x < 10 {
		for i := 0; i <= 3; i++ {
			fmt.Println(i)
			if x == 4 {
				break
			}
			// } else {
			// 	fmt.Println("not four")
			// }
		}
		x++
	}
	fmt.Println("done.")

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		fallthrough
	default:
		fmt.Println("this is default")
	}

}
