package main

import (
	"fmt"
	"time"
)

func main() {
	var st string
	var et string
	var tt string

	fmt.Println("enter start time")
	fmt.Scanln(&st)

	fmt.Println("enter end time")
	fmt.Scanln(&et)

	fmt.Println("enter tentative time")
	fmt.Scanln(&tt)
	// fmt.Println(st)
	// start_time := "4:04:17"
	// end_time := "6:04:56"
	// gtime := "5:45:59"

	stime, _ := time.Parse("15:04:05", st)
	etime, _ := time.Parse("15:04:05", et)
	ttime, _ := time.Parse("15:04:05", tt)

	if ttime.After(stime) && ttime.Before(etime) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
