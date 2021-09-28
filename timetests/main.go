package main

import (
	"fmt"
	"time"
)

func main(){
	var start string
	var end string
	var ctime string
	fmt.Println("Enter start time(hh:mm:ss)")
	fmt.Scan(&start)
	fmt.Println("enter end time(hh:mm:ss)")
	fmt.Scan(&end)
	fmt.Println("enter time to check (hh:mm:ss) ")
	fmt.Scan(&ctime)
	
	newLayout := "15:04:05"
	
	cc,_ := time.Parse(newLayout, ctime)
	
	ss, _ := time.Parse(newLayout, start)
	ee, _ := time.Parse(newLayout, end)
	ans:=checkrange(ss,ee,cc)
	
	if ans{
		fmt.Println("the given time is in range")
	}else{
		fmt.Println("time not in range")
	}

}

func checkrange(start, end, check time.Time) bool {
	if start.Before(end) {
		return check.After(start) && check.Before(end)
	}
	
	return start.Before(check) || end.After(check)
}
