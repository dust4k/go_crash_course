package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Declare and assign
	carArr := [2]string{"BMW", "Toyota"}

	fmt.Println(carArr)

	// Slices
	colorSlice := []string{"red", "blue", "pink", "yellow"}
	fmt.Println(colorSlice, len(colorSlice))
	fmt.Println(colorSlice[1:3])
}
