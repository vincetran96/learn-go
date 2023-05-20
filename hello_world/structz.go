package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Change the struct using value
func changeStruct(s Vertex) {
	s.X = s.X + 100
}

// Change the struct using pointer
func changeStruct_ptr(s *Vertex) {
	s.X = s.X + 100
}

func main() {
	// Change value
	v := Vertex{1, 2}
	changeStruct(v)
	fmt.Println("v after changeStruct:", v)
	v.X = 4
	fmt.Println("v after changing value directly:", v)
	changeStruct_ptr(&v)
	fmt.Println("v after changeStruct_ptr:", v)

	// Multiple struct variables
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}  // Y defaults to 0
		v3 = Vertex{}      // X, Y default to 0
		v4 = &Vertex{1, 2} // type *Vertex
	)
	fmt.Println(v1, v2, v3, v4)

	b := false
	if !b {
		fmt.Println("not b")
	}
}
