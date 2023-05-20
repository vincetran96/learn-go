package main

import (
	"fmt"
	"strconv"
)



func Convert(number int) string {
	result := ""
	length := strconv.Itoa(len(strconv.Itoa(number)))
	keys := []int {3, 5, 7}
	map_ := map[int]string {3: "Pling", 5: "Plang", 7: "Plong"}
	for _, key := range keys {
		if number % key == 0 {
			result += map_[key]
		}
	}
	if result == "" {
		result = length
	}
	return result
}

func main() {
	fmt.Println(Convert(210))
	fmt.Println(Convert(45))
	fmt.Println(Convert(35))
	fmt.Println(Convert(1))
}