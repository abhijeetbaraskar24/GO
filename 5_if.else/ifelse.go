package main

import "fmt"

func main() {

	// age := 23

	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// } else {
	// 	fmt.Println("not an adult")
	// }

	// if age > 18 {
	// 	fmt.Println("a")
	// } else if age > 12 {
	// 	fmt.Println("b")
	// } else {
	// 	fmt.Println("c")
	// }

	//we can declare var inside the if construct
	if a := 21; a > 12 {
		fmt.Println("is an adult")
	}

}

//go does not have terneary, you will have to use normal if else
