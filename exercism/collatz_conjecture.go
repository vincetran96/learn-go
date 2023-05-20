package main

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}
	n_steps := 0
	for {
		fmt.Printf("Step %d: n = %d\n", n_steps, n)
		if n == 1 {
			return n_steps, nil
		} else {
			n_steps += 1
			if n % 2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
		}
	}
}

func CollatzConjecture_r(n int) int {
	if n == 1 {
		return 0
	} else {
		if n % 2 == 0 {
			return 1 + CollatzConjecture_r(n / 2)
		} else {
			return 1 + CollatzConjecture_r(3*n + 1)
		}
	}
}

func main() {
	fmt.Println(CollatzConjecture(12))
	fmt.Println(CollatzConjecture(9))
	fmt.Println(CollatzConjecture(37))
	fmt.Println(CollatzConjecture(3))
	fmt.Println(CollatzConjecture_r(12))
}