package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Consume Json data in Golang")
	decodeJson()

}

func decodeJson() {

	var courseData course
	jsonData := []byte(`
		{
			"coursename": "ReactJs",
			"price": 599,
			"platform": "kamy.dev.tech",
			"tags": ["javascript","dev"]
	}
`)
	// Check for validation of json data
	validation := json.Valid(jsonData)
	fmt.Println(validation)


	json.Unmarshal(jsonData, &courseData)
	fmt.Println(courseData)

	// Some cases where you want to add key value pair

	var onlineData map[string]interface{}
	//here the use of interface is the json data can be also a int or other datatype so we use interface in it 
	json.Unmarshal(jsonData,&onlineData)
	fmt.Printf("%#v\n",onlineData)

	// i do for loop to print the key value and its type 

	for k,v := range onlineData{
		fmt.Printf("Key is :%v and value is :%v and type of data:%T\n",k,v,v)
	}
}
