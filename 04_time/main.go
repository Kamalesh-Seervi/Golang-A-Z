package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")
	mytime := time.Now()
	fmt.Println(mytime)

	fmt.Println(mytime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2002, time.December,20,20,20,0,0,time.Local)
	fmt.Println(createdDate)
	

}
