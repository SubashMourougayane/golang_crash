package main

import "fmt"

var name string = "Subash"

func main() {

	// name := "subash"
	// age := 24

	name, age := "subash", 24

	isCool := true

	fmt.Println(name, age, isCool)

	var pi float32 = 3.14
	pi = 3.145
	fmt.Printf("%T \n", pi)

}
