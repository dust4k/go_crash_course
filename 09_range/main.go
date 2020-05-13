package main

import "fmt"

func main() {
	ids := []int{33, 22, 14, 541, 1}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Loop through ids not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with maps
	ages := map[string]int{"Bob": 13, "Darren": 11, "Sue": 29, "Barb": 58}
	for k, v := range ages {
		fmt.Printf("%s: %d\n", k, v)
	}

	for k := range(ages) {
		fmt.Println("Name: " + k)
	}

}
