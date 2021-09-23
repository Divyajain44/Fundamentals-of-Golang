package main

import (
	"fmt"
)

func main() {
	//ARRAYS
	var y [5]int
	fmt.Println(y)
	y[1] = 42
	y[2] = 7
	fmt.Println(y)
	fmt.Println(len(y))

	//SLICE
	// type{value} COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)
	}

	//slicing a slice
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	//appending in slice
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)
	z := []int{234, 456, 678, 987}
	x = append(x, z...)
	fmt.Println(x)

	//deleting from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	m := make([]int, 10, 12)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	m[0] = 42
	m[9] = 999

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	m = append(m, 3423)

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	m = append(m, 3424)

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	m = append(m, 3425)

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	//MAP
	ms := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(ms["James"])
	delete(ms, "James")
	println(ms)

	fmt.Println(ms["Barnabas"])

	v, ok := ms["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := ms["Barnabas"]; ok {
		fmt.Println(v)
	}
	ms = nil // clears map
}
