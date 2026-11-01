package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var names = []string{"Kamy", "Kamy1", "kamy2"}
	fmt.Println("Value of names", names)
	fmt.Printf("Type of datatype:%T\n", names)

	names = append(names, "Kamy3", "Kamy4")
	fmt.Println("Appended Data:", names)

	names = append(names[1:])
	fmt.Println(names)

	names = append(names[1:3])
	fmt.Println(names)

	nums := make([]int, 3)
	nums[0] = 100
	nums[1] = 200
	nums[2] = 300
	// nums[3]=400
	fmt.Println("Value of nums:", nums)
	// exits and throws a runtime error
	//To avoif this we can use append to add more elements so it stores memory dynamically

	nums = append(nums, 350, 650, 250)
	fmt.Println("Value of nums:", nums)

	//Some built in methods which are avaiable for for slices but not arrays

	sort.Ints(nums)
	fmt.Println("Value of nums:", nums)
	sort_nums := sort.IntsAreSorted(nums)
	fmt.Println("Value of nums:", sort_nums)

	//Removing data by index
	

	var frds = []string{"James","Nolan","Patrick","Jason","Mrunal"}
	var index = 2
	frds = append(frds[:index],frds[index+1:]...)
	// ... you will learn in upcoming codes
	fmt.Println("Value of frds:",frds)

}
