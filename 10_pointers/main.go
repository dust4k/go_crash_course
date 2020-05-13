package main

import "fmt"

func main() {
	a := 5
	b := &a // b is a pointer to a

	fmt.Println(a, b)
	fmt.Printf("%T %T\n", a, b)

	fmt.Println(*b) // use * to read from an address that is pointed to

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}