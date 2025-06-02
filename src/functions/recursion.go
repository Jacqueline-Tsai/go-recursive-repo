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

// RecursivePower calculates the power of a number x raised to n recursively.
func RecursivePower(x, n int) int {
    if n == 0 {
        return 1
    }
    return x * RecursivePower(x, n-1)
}

// RecursiveGCD calculates the greatest common divisor of two numbers a and b recursively.
func RecursiveGCD(a, b int) int {
    if b == 0 {
        return a
    }
    return RecursiveGCD(b, a%b)
}

// CountDown prints numbers from n to 1 recursively.
func CountDown(n int) {
    if n <= 0 {
        return
    }
    fmt.Println(n)
    CountDown(n - 1)
}

// CountUp prints numbers from 1 to n recursively.
func CountUp(n int) {
    if n <= 0 {
        return
    }
    CountUp(n - 1)
    fmt.Println(n)
}

// ReverseString reverses a string recursively.
func ReverseString(s string) string {
    if len(s) == 0 {
        return s
    }
    return string(s[len(s)-1]) + ReverseString(s[:len(s)-1])
}

// Palindrome checks if a string is a palindrome recursively.
func Palindrome(s string) bool {
    if len(s) < 2 {
        return true
    }
    if s[0] != s[len(s)-1] {
        return false
    }
    return Palindrome(s[1 : len(s)-1])
}

// SumDigits calculates the sum of the digits of a number recursively.
func SumDigits(n int) int {
    if n == 0 {
        return 0
    }
    return n%10 + SumDigits(n/10)
}

// CountVowels counts the number of vowels in a string recursively.
func CountVowels(s string) int {
    if len(s) == 0 {
        return 0
    }
    count := 0
    if isVowel(s[0]) {
        count = 1
    }
    return count + CountVowels(s[1:])
}

// Helper function to check if a character is a vowel.
func isVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
           c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}