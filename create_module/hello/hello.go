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

	message, error := greetings.Hello("")
	if error != nil {
		log.Fatal(error)
	}
	
	fmt.Println(message)
	fmt.Println(pkg1.Mult_add(10,1,2,3))
	fmt.Println(pkg2.Factorial(10))

	fmt.Println("\nEnd main in hello module")
}