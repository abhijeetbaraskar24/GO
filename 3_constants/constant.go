package main

import "fmt"

//once u assign value u cant change it

//can declare the const and the var outside the func but not with this   :=

const city = "pune"

var age = 30
var name string = "golang"

func main() {

	//const name = "golang"

	//const name string= "golang"

	fmt.Println(city)

	const (
		a   = 345
		prt = "hi"
		//this is grouping const
		//two or more const value in one block
	)

}
