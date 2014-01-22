package main

import "fmt"

func fib(n uint) uint {
	// check if n == 0 or n == 1
  if n == 0 || n == 1 {
  	return n
  } else {
  	return fib(n-1) + fib(n-2)
  }
}

func main() {
	// insert any value of n within the parantheses of fib() below
	 fmt.Println(fib(10))
	// if you want a generator, comment out the above section and drop the comment below
	//var n uint
	//for n = 1; n < 10; n++ {
	//	fmt.Println(fib(n))
	//}
}

