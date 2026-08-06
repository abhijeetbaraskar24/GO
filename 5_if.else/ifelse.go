package main

import "fmt"

func main() {

	age := 23

	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// } else {
	// 	fmt.Println("not an adult")
	// }

	if age > 18 {
		fmt.Println("a")
	} else if age > 12 {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}

}
