package main

import (
	"fmt"
)

func swap(firstPtr, secondPtr *int) {
	*firstPtr, *secondPtr = *secondPtr, *firstPtr
}

func main() {
	x, y := 1 ,2
	fmt.Println(x, y)
	swap (&x, &y)
	fmt.Println(x, y)
}