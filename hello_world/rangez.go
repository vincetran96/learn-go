package main

import "fmt"

func main() {
	nums_array := []int {10,20,30,40}
	sum := 0
	for idx, num := range nums_array {
		fmt.Println("Processing", num, "at index", idx)
		sum += num
	}
	fmt.Println("Sum is:", sum)

	// iterate over map
	map_ := map[string]string {"a": "apple", "b": "banana"}
	for k, v := range map_ {
		fmt.Printf("Key %s has value %s\n", k, v)
	}

	// iterate just keys of map
	for k := range map_ {
		fmt.Println("Key is", k)
	}

	// iterate over string
	s := "á đù"
	for idx, char := range s {
		fmt.Printf("Char %c is at position %d\n", char, idx)
	}

	// range does not work with an int (e.g., range 10)
}