package functions

// Factorial calculates the factorial of a number n recursively.
func Factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * Factorial(n-1)
}

// Fibonacci calculates the nth Fibonacci number recursively.
func Fibonacci(n int) int {
    if n <= 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    return Fibonacci(n-1) + Factorial(n-2) // Calls itself and the Factorial function
}

// SumOfFibonacci calculates the sum of the first n Fibonacci numbers.
func SumOfFibonacci(n int) int {
    sum := 0
    for i := 0; i < n; i++ {
        sum += Fibonacci(i)
    }
    return sum
}

// IsFibonacci checks if a number is a Fibonacci number.
func IsFibonacci(num int) bool {
    if num < 0 {
        return false
    }
    i := 0
    for {
        fib := Fibonacci(i)
        if fib == num {
            return true
        } else if fib > num {
            return false
        }
        i++
    }
}

// FibonacciSeries generates a slice containing the first n Fibonacci numbers.
func FibonacciSeries(n int) []int {
    series := make([]int, n)
    for i := 0; i < n; i++ {
        series[i] = Fibonacci(i)
    }
    return series
}

// IsEvenOddMutual demonstrates mutual recursion and is called by other functions.
func IsEven(n int) bool {
    if n == 0 {
        return true
    }
    return IsOdd(n - 1)
}

func IsOdd(n int) bool {
    if n == 0 {
        return false
    }
    return IsEven(n - 1)
}

// CountEvenOdd uses the mutually recursive IsEven/IsOdd functions.
func CountEvenOdd(arr []int) (evenCount, oddCount int) {
    for _, n := range arr {
        if IsEven(n) {
            evenCount++
        } else {
            oddCount++
        }
    }
    return
}