package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	//composite literal func (receiver) identifier(parameters) (returns) {code}
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	p1 := person{
		first: "divya",
		last:  "Jain",
	}

	p1.met()

	//anonyms functions
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)
}

func (p person) met() {
	fmt.Println(p.first, p.last)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (a string, b bool) {
	a = fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b = true
	return a, b
}
