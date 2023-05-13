package main

import "fmt"

func arrays() {
	var a [5]int
	fmt.Println("a array:", a)

	a = [5]int{1, 2, 3, 4}
	fmt.Println("a array:", a)

	a[3] = 100
	fmt.Println("a array:", a, "with length", len(a))

	b := [5]int{10, 20, 30} // the rest is filled with 0
	fmt.Println("b array:", b)

	// Two-dimensional array (2D)
	var two_d [2][3]int
	for i := 0; i < 2; i++ {
		two_d[i] = [3]int{i + 1, i + 2, i + 3}
	}
	fmt.Println("2D array:", two_d)
}

func slices() {
	s := make([]string, 3)
	fmt.Println("s:", s, "cap s:", cap(s))
	s = []string{"a", "b", "c", "d"}
	fmt.Println("s:", s, "cap s:", cap(s))
	s = append(s, []string{"z", "x"}...)
	fmt.Println("s:", s, "cap s:", cap(s))

	c := make([]string, 1)
	copy(c, s)
	fmt.Println("c:", c, "(copied from a larger-sized slice")
	c = []string{"q", "w", "e", "r"}
	fmt.Println("c:", c, "cap c:", cap(c))

	// Two-dimensional slice
	two_d := make([][]int, 4)
	for i := 0; i < 4; i++ {
		inner_slice := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			inner_slice[j] = j
		}
		two_d[i] = inner_slice
	}
	fmt.Println("2D slice:", two_d)
}

func main() {
	// arrays()
	slices()
}
