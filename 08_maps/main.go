package main

import "fmt"

func main() {
	// Define a map
	emails := make(map[string]string)
	ages := map[string]int{"Bob":13, "Darren":11}

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Darren"] = "darren@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))
	fmt.Println(ages)

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}