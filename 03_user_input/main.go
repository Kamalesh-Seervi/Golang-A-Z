package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "hi Users"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza:")

	// ReadString returns (string, error). The "comma ok / err" idiom lets us check the error. 
	// Go has no try/catch: errors are returned as values. The _ below discards the error. 

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("type of input: %T",input)

	// ReadString always returns a string, even when the user types digits.
	// The value also keeps the trailing newline.

}
