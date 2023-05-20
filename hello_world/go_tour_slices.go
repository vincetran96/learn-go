package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/tour/pic"
)

var rand_source = rand.New(rand.NewSource(time.Now().UnixNano()))

func Pic(dx, dy int) [][]uint8 {
	max_num := 255
	result := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		result[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			result[i][j] = uint8(rand_source.Intn(max_num))
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}