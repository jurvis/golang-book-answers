package main

import "fmt"

func main() {
    fmt.Print("Enter a length in feet: ")
    var input float64
    fmt.Scanf("%f", &input)

    // use formula as given
    output := input * 0.3048

    //print the output from above operation
    fmt.Println("The length in meters is: ", output)
}