package main

import (
	"fmt"
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	wordFormatted := nonAlphanumericRegex.ReplaceAllString(word, "")
	charCount := make(map[string]int)
	for _, char := range(wordFormatted) {
		charStr := strings.ToLower(string(char))
		_, isIn := charCount[charStr]
		if !isIn {
			charCount[charStr] = 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsIsogram("six-year-old"))
	fmt.Println(IsIsogram("isograms"))
	fmt.Println(IsIsogram("Aapple"))
	fmt.Println(IsIsogram("   "))
}