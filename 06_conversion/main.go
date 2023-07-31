package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:",input)

	// now if i try to add +1 to the input it shows err because 4 acts as a string 
	// to overcome this we need to use
	// strconv

	numadd, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	// we use trimspace because if i dont use it, it takes 4\n together which is issue in strconv
	// trim space removes the spaces so the final numadd will be only 4
	if err!= nil{
		fmt.Println(err)
} else{
	newRating := numadd +1
	fmt.Println("New rating:",newRating)
}
}