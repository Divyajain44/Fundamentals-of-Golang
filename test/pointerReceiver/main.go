package main

import "fmt"

type person struct {
	First string
}

type human interface {
	speak()
}

func (*person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		First: "divya",
	}
	saySomething(&p1)
}
