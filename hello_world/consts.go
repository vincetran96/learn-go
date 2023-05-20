package main

import (
	"fmt"
	"math"
)

const s string = "const string"

func main() {
	fmt.Println(s)

	const n float64 = 500000000

	const d = 3e20 / n

	const u, v float32 = 0, 3

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}