package main

import "fmt"

// for - > only construct in go for looping

func main() {

	//while loop

	//i := 0
	//for i <= 3 {
	//	fmt.Println(i)
	//	i += 1
	//}

	//classic for loop

	// for i := 1; i <= 3; i++ {

	// 	if i == 1 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }

	//go 1.22 new verison have feature of range

	for i := range 4 {
		fmt.Println(i)
	}
}
