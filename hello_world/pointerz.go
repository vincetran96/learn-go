package main

import "fmt"

func add(n int, a int) {
	n += a
}

func add_ptr(n *int, a int) {
	*n += a
}

func main() {
	x, y := 10, 20

	var ptr_x *int
	ptr_x = &x
	ptr_y := &y

	fmt.Println("ptr_x value:", *ptr_x)
	fmt.Println("ptr_y value:", *ptr_y)

	*ptr_x = 100
	fmt.Println("new ptr_x value:", *ptr_x)
	fmt.Println("x value after changing ptr_x:", x)
	x = 500
	fmt.Println("new x value:", x)
	fmt.Println("ptr_x value after changing x:", *ptr_x)

	add(x, 10)
	fmt.Println("x value after calling add(x, 10):", x)

}
