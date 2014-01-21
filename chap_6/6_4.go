//Write a program that finds the smallest number in this list:
//x := []int{
//    48,96,86,68,
//    57,82,63,70,
//    37,34,83,27,
//    19,97, 9,17,
//}
package main

import "fmt"

func findSmallestNumber() {
	x := []int{
	    48,96,86,68,
	    57,82,63,70,
	    37,34,83,27,
	    19,97, 9,17,
	}

	var number int = x[0]

	for i := 0; i < len(x); i+ {
		if x[i] < number {
			number=x[i]
		}
	}

	fmt.Println("lowest number in this list is: ", number)
}

func main () {
	findSmallestNumber()
}
