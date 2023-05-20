package main

import (
	"fmt"
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b must have equal length")
	}
	distance := 0
	for idx, c := range a {
		char_a, char_b := string(c), string(b[idx])
		fmt.Println("checking", char_a, "and", char_b)
		if char_a != char_b {
			distance += 1
		}
	}
	return distance, nil
}

func main() {
	fmt.Println(Distance("zxcv", "zxcb"))
	fmt.Println(Distance("zxcv", "a"))
	fmt.Println(Distance("zxcv", "zxcv"))
	fmt.Println(Distance("", ""))
}