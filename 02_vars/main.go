package main

import "fmt"

func main() {
	// using var
	var name string = "Billy" // explicitly defining the type
	var month = "December" // don't need to specify type if it is implied
	var age int32 = 37
	var isCool = true
	const isDone = false

	fmt.Println(name)
	fmt.Println(month, age, isCool)

	// shorthand way
	city := "Toronto"
	size := 1.3 // inferred as float64

	// combining
	color, email := "red", "me@abc.com"

	fmt.Printf("%T %T %T\n", age, month, isDone) // useful for determining the type
	fmt.Println(city, size, color, email)
}
