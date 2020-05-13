package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Sam", lastName: "Jones", city: "Toronto", gender: "F", age: 22}
	person2 := Person{"Sam", "Jones", "Toronto", "F", 22}
	fmt.Println(person1, person2)

	// Get single field
	fmt.Println(person1.firstName)

	person1.age++
	fmt.Println(person1)
	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("Vanwinkle")
	fmt.Println(person1.greet())
}
