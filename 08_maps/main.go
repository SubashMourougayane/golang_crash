package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["bob"] = "bob.gmail.com"
	emails["brad"] = "brad.gmail.com"
	emails["mike"] = "mike.gmail.com"

	fmt.Println(emails)

	fmt.Println(len(emails))

	delete(emails, "bob")
	fmt.Println(len(emails))

}
