package main

import "fmt"

func Reverse(s string) string {
	var s_rev string
	s_utf := []rune(s)
	for i := len(s_utf)-1; i >= 0; i-- {
		fmt.Println(string(s[i]))
		s_rev += string(s_utf[i])
	}
	return s_rev
}

func main() {
	fmt.Println(Reverse("Hello, 世界"))	
}