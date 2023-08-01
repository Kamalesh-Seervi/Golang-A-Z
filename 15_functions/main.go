package main

import "fmt"

func greet() {
	fmt.Println("Welcome of Kamy GoLang")
}


//Using arguments
func adder(num1 int, num2 int) int {  //this is called function signature because we mention our final return to be int we add int at the final of the func
	return num1 + num2
}

func proAdder(values ...int) (int,string){
	total := 0
	for _,val := range values{
		total+=val
	}
	return total,"hi guys this to show we can add strings too"
}
func main() {
	fmt.Println("Fucntions in Golang")
	greet()
	result:=adder(2,5)
	fmt.Println(result)

	proRes,Stringop:=proAdder(2,5,3,2,65)
	fmt.Println("Final proResult:",proRes)
	fmt.Println(Stringop)

}
