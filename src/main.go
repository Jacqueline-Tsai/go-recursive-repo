package main

import (
    "fmt"
    "go-recursive-repo/src/functions"
)

func main() {
    // Example usage of math utility functions
    sum := functions.Add(5, 3)
    fmt.Println("Sum:", sum)

    difference := functions.Subtract(10, 4)
    fmt.Println("Difference:", difference)

    product := functions.Multiply(7, 6)
    fmt.Println("Product:", product)

    quotient := functions.Divide(20, 4)
    fmt.Println("Quotient:", quotient)

    // Example usage of string utility functions
    concatenated := functions.Concatenate("Hello, ", "World!")
    fmt.Println("Concatenated String:", concatenated)

    splitString := functions.Split("Go is awesome", " ")
    fmt.Println("Split String:", splitString)

    reversed := functions.Reverse("Golang")
    fmt.Println("Reversed String:", reversed)

    // Example usage of recursive functions
    factorial := functions.Factorial(5)
    fmt.Println("Factorial of 5:", factorial)

    fibonacci := functions.Fibonacci(7)
    fmt.Println("Fibonacci of 7:", fibonacci)
}