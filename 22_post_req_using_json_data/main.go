package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post req using json data in Golang")
	postReq()
}

func postReq() {
	const myUrl = "http://localhost:1111/post"

	// we create a fake payload
	jsondata := strings.NewReader(`
	{
		"name":"Kamalesh",
		"age":25,
		"place:"Chennai"
	}
`)
	response, _ := http.Post(myUrl, "application/json", jsondata)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
	defer response.Body.Close()

}
