package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")
	now := time.Now()
	fmt.Println(now)

	// Go formats dates using the reference time Mon Jan 2 15:04:05 2006 (01-02-2006 = month-day-year).
	fmt.Println(now.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2002, time.December, 20, 20, 20, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02 Jan 2006"))

	// Since returns a time.Duration measured from the given time until now.
	fmt.Println("Hours since created:", int(time.Since(createdDate).Hours()))
}
