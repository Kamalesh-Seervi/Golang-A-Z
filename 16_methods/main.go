package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	//No inheritance and No super and parent

	Kamalesh := User{"Kamy",21,"kamy@06gmail.com",true}
	fmt.Println("Userdata:",Kamalesh)
	fmt.Printf("Userdata %+v\n",Kamalesh)
	fmt.Println("Userdata",Kamalesh.Name,Kamalesh.Email)
	Kamalesh.GetStatus()
	Kamalesh.NewMail()
	

}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus(){
	fmt.Println(u.Status)
}

func (u User) NewMail(){
	u.Email="test@06gmail.com"
	fmt.Println("New Email:",u.Email)
}