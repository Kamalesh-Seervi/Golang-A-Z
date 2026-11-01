package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi users")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza (1-5):")

	// ReadString returns (string, error). The "comma ok / err" idiom lets us check the error.
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Could not read input:", err)
		return
	}

	// ReadString always returns a string, even when the user types digits.
	// The value keeps the trailing newline, so trim it before converting.
	input = strings.TrimSpace(input)
	fmt.Printf("type of input: %T\n", input)

	// Convert the string to a number with strconv.ParseFloat.
	rating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Rating must be a number:", err)
		return
	}

	fmt.Println("Thanks for rating,", rating)
	fmt.Printf("type of rating: %T\n", rating)
	fmt.Println("Rating plus 1:", rating+1)
}
