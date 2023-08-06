package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int	`json:"price"`
	Platform string `json:"platform"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Create Json data in Golang")
	encodedJson()
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func encodedJson() {
	kamycourses := []course{
		{"ReactJs", 599, "kamy.dev.tech", "abc123", []string{"javascript", "dev"}},
		{"MERN Stack", 1099, "kamy.dev.tech", "kamy123", []string{"Full-stack", "dev"}},
		{"NextJs", 899, "kamy.dev.tech", "next123",nil},
	}
	jsonData, err := json.MarshalIndent(kamycourses, "", "\t")
	checkNilError(err)
	fmt.Println(string(jsonData))
}
