package main

import (
	"fmt"
)

// write a function called half that accepts an integer and returns an integer and string
func half (x int) (int, string) {
	// divide input by 2
	output := x / 2

	// check if output is multiple of 2 - if yes, input is even.
	if output % 2 == 0 {
		return output , "true"
	} else {
		return output , "false"
	}
}

func main() {
	// assign returned values to "x" and "y"
	// you can enter any integer between the parantheses of half()
	x, y := half(2)
	//print result
	fmt.Printf("(%d, %s)\n", x, y)
}

