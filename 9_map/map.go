package main

import "fmt"

// MAP -> HASH,OBJ,DICT

func main() {

	//creating map

	m := make(map[string]string)

	//setting value to it

	m["name"] = "golang"
	m["area"] = "backend"

	//getting the value

	fmt.Println(m["name"], m["area"])

	//IMP if the key does not exist then the println of the key is is going to be empty for string, zero for int and false for the boolean

	ma := make(map[string]int)

	ma["age"] = 30
	ma["price"] = 40

	fmt.Println(ma)

	//getting the len of the map or say no of key value pairs in it
	fmt.Println(len(m))

	//delete the key and value pair
	delete(ma, "price")

	fmt.Println(ma)

	//to delete all the pairs at a time or say clear the map

	clear(ma)
	fmt.Println(ma)
}
