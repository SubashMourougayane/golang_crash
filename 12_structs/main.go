package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am  " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "F", age: 25}

	person2 := Person{"Helly", "Keller", "Boston", "F", 25}

	fmt.Println(person1.firstName)

	person1.age++
	fmt.Println(person1, person2)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")

	fmt.Println(person1.greet())

}
