package main

import "fmt"

func main() {
	fmt.Println("IF else in GoLang")

	//use open curly brace just after the if condition dont put in next line it will throw error becuase of lexers
	num := 35
	var result string
	if num > 50 {
		result = "Good"
	} else {
		result = "Bad"
	}
	fmt.Println(result)

	// another way if useful when you use to create sites
	// get user imput in if directly

	if nums := 3; nums < 10 {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}

	
}
