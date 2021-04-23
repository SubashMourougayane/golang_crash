package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, *b)
	fmt.Printf("%T , %T \n", a, b)

	*b = 10
	fmt.Println(a)

}
