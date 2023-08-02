package main

import (
	"fmt"

)

// defers are like last in first out 
func main() {
	fmt.Println("Defer in Golang")

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("\nthree")
	fmt.Println("hello")
	mydefer()
}

func mydefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
		
	}
}
