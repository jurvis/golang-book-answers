package main

import "fmt"

func main() {
	// declare variable "number" as 1
    number := 1
  // create for loop for number less than or equals to 100 since question says number has to be less than 100
    for number <= 100 {
    	// check if number is a multiple of 3, if returns true, print number
        if number % 3 == 0 {
        	fmt.Println(number)
        }
      // move on to next number
      number += 1
    }
}