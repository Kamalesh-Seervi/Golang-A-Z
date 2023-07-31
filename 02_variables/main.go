package main

import "fmt"

const token int = 134325 //public

func main()  {
	// string
	var name string = "Kamalesh"
	fmt.Println("Hello, ",name)
	fmt.Printf("variable of type: %T \n",name)

	//int
	var num int = 245
	fmt.Println(num)
	fmt.Printf("variable of type: %T \n",num)

	//bool
	var case1 bool = true
	fmt.Println(case1)
	fmt.Printf("variable of type: %T \n",case1)

	//float are of two types float32 and 64

	var flt float32 = 4592.028502803085028
	fmt.Println(flt)
	fmt.Printf("variable of type: %T \n",flt)


	// default values and some alias

	// int default is zero 
	var new_num int
	fmt.Println(new_num)
	fmt.Printf("variable of type: %T \n",new_num)

	// prints nothing empty 
	var named string
	fmt.Println(named)
	fmt.Printf("variable of type: %T \n",named)

	//implicit type 

	var nametyp = "Kamalesh007"
	fmt.Println(nametyp)
	fmt.Printf("variable of type: %T \n",nametyp)

	// without var we can use walrus to define a varible .
	//But you can only do it inside a method or func out the func it throws a error

	username := "Kamy"
	fmt.Println(username)
	fmt.Printf("variable of type: %T \n",username)

	// print a const which is defined outside the func
	fmt.Println(token)
	fmt.Printf("variable of type: %T \n",token)
}