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

// Closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println(sum_mult(10,1,2,3))
	
	a, b := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(10), b(-20))
	}
}