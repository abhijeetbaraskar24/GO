package main

import "fmt"

// array is numbered seq of fixed length and it store data of same type

func main() {

	var num [4]int

	//len of array
	fmt.Println("len of array is : ", len(num), " .")

	//add element in array
	num[2] = 4

	//print whole array here as the array type is int so the rest of the elements in the array which dont have any value yet are set to zero automatically
	fmt.Println(num)

	//so when u dont give values to the array
	//it automatically gives zeroed values
	//int -> 0, string -> "" empty value, boolean -> false

	//declare value in a single line
	nums := [5]int{1, 2, 3, 4, 5}

	fmt.Println(nums)

	//2D arrays

	number := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(number)

	// - fixed size, it is predictable
	// - memory Optimization
	// - constant time access -> (cause we know the index numbers and also the size of the array )

}
