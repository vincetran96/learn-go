package greetings

import (
	"fmt"
	"errors"
	"example.com/greetings/pkg1"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	
	message := fmt.Sprintf("Hello, %v", name)
	n := pkg1.Mult_add(10,1,2,3)
	fmt.Println("Testing Mult_add in package greetings:", n)
	return message, nil
}