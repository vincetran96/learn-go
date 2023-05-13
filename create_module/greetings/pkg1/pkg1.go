package pkg1

func Mult_add(mult int, nums... int) int {
	sum := 0
	for _, num := range nums {
		sum += mult * num
	}
	return sum
}