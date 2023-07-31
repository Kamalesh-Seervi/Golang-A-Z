package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	lang := make(map[string]string)
	//map[key]value
	lang["Js"]="Javascript"
	lang["Py"]="Python"
	lang["P"]="Php"
	lang["Rb"]="Ruby"
	fmt.Println("Types of langs:",lang)
	fmt.Println("Types of langs:",lang["Py"])
	//Print a particular value


	// Deleting a Key/value
	delete(lang,"Rb")
	fmt.Println("Types of langs:",lang)

	// Loops in Maps
	for key,value := range(lang){
		fmt.Printf("Key is %v and Value is %v\n",key,value)
	}

}
