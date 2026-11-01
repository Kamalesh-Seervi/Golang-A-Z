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

	// comma ok || err 
	// similar to try catch 

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("type of input: %T",input)

	// Here i notice that it shows type of input as string but it is integer 
	// we will see this type of catch error in next folder 

}
