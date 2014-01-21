// PLEASE VIEW 7_1.go FIRST!
// This version allows user input instead of having to hard-code the value of the input.
// 7_1.go answers the question more accur

package main

import (
	"fmt"
)

// write a function called half that accepts an integer and returns an integer and string
func half () {
	fmt.Print("Enter a number: ")
	// create a floating point variable (since output is int, your input has to be int)
	var input int
	fmt.Scanf("%f", &input)
	// divide input by 2
	output := input / 2

	// check if output is multiple of 2 - if yes, input is even.
	if output % 2 == 0 {
		fmt.Printf("(%d, true)\n", output)
	} else {
		fmt.Printf("(%d, false)\n", output)
	}
}

func main() {
	half()
}

