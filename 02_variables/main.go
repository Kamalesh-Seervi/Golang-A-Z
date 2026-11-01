package main

import "fmt"

// Lowercase first letter means unexported: usable only inside this package.
// An uppercase name such as Version would be exported to other packages.
const token int = 134325

func main()  {
	// string
	var name string = "tester"
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

	// float32 holds about 7 significant digits, so the extra digits below are rounded off

	var flt float32 = 4592.028002853085028
	fmt.Println(flt)
	fmt.Printf("variable of type: %T \n",flt)


	// Zero values: a variable declared without a value gets the default for its type.

	// int default is zero 
	var new_num int
	fmt.Println(new_num)
	fmt.Printf("variable of type: %T \n",new_num)

	// string default is "" (prints as an empty line) 
	var named string
	fmt.Println(named)
	fmt.Printf("variable of type: %T \n",named)

	// Type inference: the compiler works out the type from the value. 

	var nametyp = "tester007"
	fmt.Println(nametyp)
	fmt.Printf("variable of type: %T \n",nametyp)

	// := declares and infers the type in one step.
	// It is only allowed inside a function; outside a function it is a compile error.

	username := "Kamy"
	fmt.Println(username)
	fmt.Printf("variable of type: %T \n",username)

	// Constants declared outside the function.
	fmt.Println(token)
	fmt.Printf("variable of type: %T \n",token)
}