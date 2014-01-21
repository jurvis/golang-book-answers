package main

import "fmt"

func main() {
    fmt.Print("Enter a temperature in Fahrenheit: ")
    // create a floating point variable (since output is float64, your input has to be float64)
    var input float64
    fmt.Scanf("%f", &input)

    // use formula as given and create a variable named output
    output := (input - 32) * 5/9

    //print the output from above operation
    fmt.Println("The temperature in Celsius is: ", output)
}