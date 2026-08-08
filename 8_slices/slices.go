package main

import (
	"fmt"
	"slices"
)

// slice -> it is dynamic (when u dont the fixed size)
// most used construct in go
// + useful methods

func main() {
	// //how to declare slice

	// var num []int
	// //uninitialized slice is nill means null
	// fmt.Println(num)
	// fmt.Println(num == nil)
	// fmt.Println(len(num))

	// //slice with make
	// var nos = make([]int, 2)
	// fmt.Println(nos)

	// var nums = make([]int, 3, 5)
	// // ([]int, 0, 5) so ur added elements start from 0
	// //the 3 is len and 5 is capacity and they resize as we addd elements

	// fmt.Println(nums)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4) //add elements in the slice

	// fmt.Println(nums)

	// fmt.Println(cap(nums)) //shows the cap of the slice after we add element more than the given cap and it resizes itself
	// fmt.Println(len(nums))

	// // above other ways

	// var numus = make([]int, 0, 5) // ([]int, 0, 5) so ur added elements start from 0
	// fmt.Println(numus)
	// numus = append(numus, 1)
	// numus = append(numus, 2)
	// fmt.Println(numus)

	// //direct creation

	// noms := []int{}
	// fmt.Println(noms)
	// fmt.Println(len(noms))
	// fmt.Println(cap(noms))

	// //copy function

	// var nums2 = make([]int, len(numus))
	// copy(nums2, numus)
	// fmt.Println(numus, nums2)

	//slice operator

	// var nums = []int{1, 2, 3}
	// fmt.Println(nums)
	// fmt.Println(nums[0:1])

	//slice package

	var nums1 = []int{}
	var nums2 = []int{}

	fmt.Println(slices.Equal(nums1, nums2))

	//2D slices

	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
}
