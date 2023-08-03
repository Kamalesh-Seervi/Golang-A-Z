package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com"

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Web Reqs in Golang")

	response, err := http.Get(url)
	checkNilError(err)
	//Get request
	fmt.Println(response)


	defer response.Body.Close()
//you need to close the connectuon of call otherwise it will be always open state


	datatype,err:=io.ReadAll(response.Body)
	checkNilError(err)

	fmt.Println("Get response content:\n",string(datatype))

}
