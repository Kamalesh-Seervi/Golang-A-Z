package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Routines in Golang")
	go greet("hello")
	greet("world")

// After adding go to hello it calls the thread but the receiving of it takes time by then it prints world.
// so to overcome this we can add some delay to the for loop so we can get callback of the hello in middle
// this is the not right way in next code we will learn how to do correct way
}

func greet(s string)  {
	for i := 0; i <5; i++ {
		time.Sleep(3*time.Millisecond)
		fmt.Println(s)
	}
}