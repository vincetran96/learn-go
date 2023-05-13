package main

import "fmt"

// Multiply and sum
func sum_mult(mult int, nums... int) int {
	sum := 0
	for _, num := range nums {
		sum += mult * num
	}
	return sum
}

func main() {
	fmt.Println(sum_mult(10,1,2,3))
}