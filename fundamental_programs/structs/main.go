package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

//embedded struct
type SecretAgent struct {
	person
	ltk bool
}

func main() {
	Sa := SecretAgent{
		person: person{
			first: "Divya",
			last:  "Jain",
			age:   21,
		},
		ltk: true,
	}

	p2 := person{
		first: "Ayush",
		last:  "Goyal",
		age:   22,
	}

	fmt.Println(Sa)
	fmt.Println(p2)

	fmt.Println(Sa.first, p2.last, p2.age)

}
