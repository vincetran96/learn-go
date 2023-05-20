package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"example.com/greetings/pkg1"
	"example.com/greetings/pkg2"
)

func main() {
	fmt.Println("Start main in hello module\n")

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, error := greetings.Hello("John")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
	fmt.Println(pkg1.Mult_add(10, 1, 2, 3))
	fmt.Println(pkg2.Factorial(10))

	messages, error_1 := greetings.Hellos([]string{"John", "Kevin"})
	if error_1 != nil {
		log.Fatal(error_1)
	}
	fmt.Println(messages)

	fmt.Println("\nEnd main in hello module")
}
