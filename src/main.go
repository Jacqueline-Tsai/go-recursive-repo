package main

import (
    "fmt"
    "go-recursive-repo/src/functions"
)

func main() {
    // Example usage of recursive functions
    factorial := functions.Factorial(5)
    fmt.Println("Factorial of 5:", factorial)

    fibonacci := functions.Fibonacci(7)
    fmt.Println("Fibonacci of 7:", fibonacci)

    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evenCount, oddCount := functions.CountEvenOdd(arr)
    fmt.Printf("Even count: %d, Odd count: %d\n", evenCount, oddCount)
}