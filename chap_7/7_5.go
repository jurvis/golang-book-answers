package main

import "fmt"

func fib() func() int {
    x, y := 0, 1
    return func() (r int) {
        r = x
        x, y = y, x + y
        return
    }
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}