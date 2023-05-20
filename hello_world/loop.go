package main

import (
	"fmt"
	"time"
	"reflect"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func func_if (num int) {
	// Loop from 0 to num and do some stuff
	for i := 0; i <= num; i++ {
		// conditions
		if_div_3 := i%3 == 0
		if_special := (i%2 == 0 && i > 4) || (i%5 == 0)

		if if_div_3 {
			msg := fmt.Sprintf("Found %v divisible by 3\n", i)
			fmt.Print(msg)
		} else if if_special {
			fmt.Println("Found some special number:", i)
		}
	}
}

func func_switch(num int) {
	// Write numbers from 0 to 3 as text, ignore the rest
	for i := 0; i <= num; i++ {
		fmt.Print("Write ", i, " as: ")
		switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		default: fmt.Println("Nah, forget it")
		}
	}

	// Switch without using an explicit variable
	// Multiple expressions in one case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	default: {
		fmt.Println("Weekday :(")
	}
	}

	// Switch without explicit expression = express if/else logic
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Before noon")
	default: 
		fmt.Println("After noon =)")
	}

	findType := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("Bool")
		case int:
			fmt.Println("Int")
		default:
			fmt.Println("Something else")
		}
	}
	findType(true)
	findType(9)
}

func main() {
	// func_if(10)
	func_switch(10)
	fmt.Println(factorial(3))
}