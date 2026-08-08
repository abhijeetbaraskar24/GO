package main

import (
	"fmt"
	"maps"
)

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

	// clear(ma)
	// fmt.Println(ma)

	mos := map[string]int{"price": 40, "phone": 3}
	fmt.Println(mos)

	//check whether the element is present in the map or not

	v, ok := mos["phone"]
	//v is the value retrieved from the map for key "phone" . if the key is missing , v is the zero value
	fmt.Println(v)
	//ok is a boolean that is true when the key exists in the map and false when it does not
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not okay")
	}

	//to compare two diff maps we are going to use equal from maps named package

	map1 := map[string]int{"price": 40, "phone": 3}
	map2 := map[string]int{"price": 40, "phone": 3}

	fmt.Println(maps.Equal(map1, map2))
}
