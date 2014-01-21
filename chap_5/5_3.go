package main

import "fmt"

func main() {
	// declare variable "number" as 1
    number := 1
  // create for loop for number less than or equals to 100 since question says number has to be less than 100
    for number <= 100 {
    	// check if number is a multiple of 15 (for values that are multiples of both 3 and 5), if returns true,
    	// print "FizzBuzz"
        if number % 15 == 0 {
        	fmt.Println("FizzBuzz")
        } else if number % 3 == 0 {
        	fmt.Println("Fizz")
        } else if number % 5 == 0 {
        	fmt.Println("Buzz")
        } else {
        	fmt.Println(number)
        }
      // move on to next number
      number += 1
    }
}