package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	//No inheritance and No super and parent

	Kamalesh := User{"Kamy",21,"kamy@06gmail.com",true}
	fmt.Println("Userdata:",Kamalesh)
	fmt.Printf("Userdata %+v\n",Kamalesh)
	fmt.Println("Userdata",Kamalesh.Name,Kamalesh.Email)

}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
