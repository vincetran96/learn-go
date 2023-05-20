package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"example.com/greetings/pkg1"
)

// var rand_source rand.Rand
var rand_source = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomFormat() string {
	var formats = []string{
		"Hi %v. Welcome",
		"Hello %v!",
		"Yo, what's up %v?",
	}
	return formats[rand_source.Intn(len(formats))]
}

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(RandomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}
		messages[name] = message
	}

	n := pkg1.Mult_add(10, 1, 2, 3)
	fmt.Println("Testing Mult_add in package greetings:", n)
	return messages, nil
}
