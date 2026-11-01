package main

import "fmt"

// Package-level constant. The lowercase first letter makes it unexported:
// it is visible only inside this package.
const token int = 134325

// Package-level constant. The uppercase first letter makes it exported:
// other packages can use it as variables.Version.
const Version string = "1.0"

func main() {
	// string
	var name string = "tester"
	fmt.Println("Hello,", name)
	fmt.Printf("variable of type: %T\n", name)

	// int
	var num int = 245
	fmt.Println(num)
	fmt.Printf("variable of type: %T\n", num)

	// bool
	var isActive bool = true
	fmt.Println(isActive)
	fmt.Printf("variable of type: %T\n", isActive)

	// float32 holds about 7 significant digits, so the extra digits below are rounded off
	var price float32 = 4592.028002853085028
	fmt.Println(price)
	fmt.Printf("variable of type: %T\n", price)

	// Zero values: a variable declared without a value gets the default for its type.
	var count int // 0
	fmt.Println(count)
	fmt.Printf("variable of type: %T\n", count)

	var empty string // "" (prints as an empty line)
	fmt.Println(empty)
	fmt.Printf("variable of type: %T\n", empty)

	// Type inference: the compiler works out the type from the value.
	var inferred = "tester007"
	fmt.Println(inferred)
	fmt.Printf("variable of type: %T\n", inferred)

	// := declares and infers in one step. It is only allowed inside a function.
	username := "Kamy"
	fmt.Println(username)
	fmt.Printf("variable of type: %T\n", username)

	// Constants declared outside the function.
	fmt.Println(token)
	fmt.Printf("variable of type: %T\n", token)
	fmt.Println("Version:", Version)
}
