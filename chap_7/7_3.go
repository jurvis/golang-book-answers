package main

import (
	"fmt"
)

// create function "Gretest Number"
func GreatestNumber(first int, rest ...int) int {
	for _, x := range rest {
		if x > first {
			first =x
		}
	}
	return first
}

func main() {
	numbers := []int{7, 4, -3, -2, 8, 9, 10, 11, 99}
	fmt.Println(GreatestNumber(numbers[0], numbers[1:]...))
}