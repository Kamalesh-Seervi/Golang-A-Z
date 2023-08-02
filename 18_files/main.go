package main

import (
	"fmt"
	"io"
	"os"
)

// Create is for -- os operation
// -- io for storing the content in specified file .
// -- ioutil is used for reading a file and some manipulation

func main() {
	fmt.Println("Files in Golang")
	content := "Hello folks this is kamy Welcome to my Golang series"
	file, err := os.Create("./Welcome.txt")
	// os.create creates a new file with specified directory
	checkNilError(err)
	length, err := io.WriteString(file, content)
	//It put the content in the text file and return the length of it .
	checkNilError(err)
	fmt.Println("Length of the Content:", length)
	defer file.Close()
	readFile("./Welcome.txt")
}

// Since ioutil go deprecated we use os.readfile
func readFile(filename string) {
	datatype, err := os.ReadFile(filename)
	// By default readfile data are in byte format we need to convert into string
	checkNilError(err)
	fmt.Println(string(datatype))

}

//best practice for error is ise function dont type if everytime

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
