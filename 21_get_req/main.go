package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Get Request in Golang")
	getReq()

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func getReq() {
	const myUrl = "http://localhost:1111/get"
	response, err := http.Get(myUrl)
	checkNilError(err)

	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

	content, err := io.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println(string(content))
	defer response.Body.Close()

	// another way to print the string(content) is by using strings.builder

	var strcontent strings.Builder
	strcontent.Write(content)
	fmt.Println(strcontent.String())

}
