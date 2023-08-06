package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Handling form data in Golang")
	postFormReq()
}

func postFormReq() {
	const myUrl = "http://localhost:1111/postform"

	// to add form data use url.Values{} method to add the data

	data:=url.Values{}
	data.Add("name","Kamalesh")
	data.Add("age","3")
	data.Add("place","Chennai")

	response,_:=http.PostForm(myUrl,data)
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))
	defer response.Body.Close()

}
