package main

import "fmt"

func main() {
	// //simple switch
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// 	//after each case in other programming lang we need to use break but in go lang we dont need it to write by urself

	// case 2:
	// 	fmt.Println("two")

	// case 3:
	// 	fmt.Println("three")

	// case 4:
	// 	fmt.Println("four")

	// 	//default is optional
	// default:
	// 	fmt.Println("other")
	// }

	//multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its holiday")
	// default:
	// 	fmt.Println("its working day")
	// }

	//type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI(true)
}
