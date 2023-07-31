package main

import "fmt"

func main() {
	fmt.Println("Loops and Goto in Golang")

	name := []string{"Kamalesh", "Mamtha", "Suresh"}
	places := []string{"Paris", "Europe", "NewYork"}

	// Normal way
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
	// Using range
	for i := range places {
		fmt.Println(places[i])
	}

	//use comma syntax
	for index, val := range places {
		fmt.Printf("index is :%v and value is %v\n", index, val)
	}

//goto

// kamy:
// 	fmt.Println("hi guys i am Kamalesh")

// 	num := 12
// 	for num > 10 {
// 		if num == 12 {
// 			goto kamy
// 		}
// 		fmt.Println("Value is :", num)
// 		num++
// 	}

// 	goto kamy
}
