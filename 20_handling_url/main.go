package main

import (
	"fmt"
	"net/url"
)

const myUrl="https://www.google.com:443/learn?courses=reactjs&"

func checkNilError(err error)  {
	if err!=nil{
		panic(err)
	}
}

func main() {
	fmt.Println("Handling url in golang")

	output,err:=url.Parse(myUrl)
	checkNilError(err)

	//type of parameter in Parse
	fmt.Println(output.Scheme)
	fmt.Println(output.Port())
	fmt.Println(output.Path)
	fmt.Println(output.Host)
	// fmt.Println(output.RawQuery)


	//we can use query method to query the url it has many porperties
	query:=output.Query()
	checkNilError(err)
	fmt.Println(query)
	fmt.Println(query["courses"])

	// We can also build url from parts

	partsofurl:=&url.URL{
		Scheme: "https",
		Host: "kamy.com",
		Path: "/learn",
		RawPath: "user=kamy",
	}
	anotherURL:=partsofurl.String()
	fmt.Println(anotherURL)

}
