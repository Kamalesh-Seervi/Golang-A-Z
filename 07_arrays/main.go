package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var names [3]string
	names[0]= "kamy"
	names[1]="kamy1"
	names[2]="kamy2"
	fmt.Println("Value of names:",names)
	fmt.Println("Value of names:",len(names))
	// Note that even if you put only elements in a array it prints the len of array as 3 (Note it)
	// it you dont add the 3 element in array then i will print a space in that index region 

	//Another way to define arrays
	var places = [3]string{"USA","UK","Russia"}
	fmt.Println("Value of the places",places)


	//In Golang we mostly dont use arrays there are other data types which are used in golang eg: maps
}
