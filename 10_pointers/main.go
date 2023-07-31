package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int
	fmt.Println("Value of the pointer is:", ptr)
	// It will return nil

	number :=20
	pointer := &number
	fmt.Println("Value of the new pointer:",pointer)
	fmt.Println("Value of the new pointer:",*pointer)

	*pointer = *pointer +5
	fmt.Println("Value of the updated pointer:",*pointer)


}
