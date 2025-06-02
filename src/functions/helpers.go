package functions

func IsEven(n int) bool {
    return n%2 == 0
}

func IsOdd(n int) bool {
    return n%2 != 0
}

func IsPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func GCD(a, b int) int {
    if b == 0 {
        return a
    }
    return GCD(b, a%b)
}

func LCM(a, b int) int {
    return (a * b) / GCD(a, b)
}

func SumOfDigits(n int) int {
    if n == 0 {
        return 0
    }
    return n%10 + SumOfDigits(n/10)
}

func ReverseString(s string) string {
    if len(s) == 0 {
        return s
    }
    return string(s[len(s)-1]) + ReverseString(s[:len(s)-1])
}

func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Factorial(n-1)
}

func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

func Power(base, exp int) int {
    if exp == 0 {
        return 1
    }
    return base * Power(base, exp-1)
}

func CountVowels(s string) int {
    if len(s) == 0 {
        return 0
    }
    vowels := "aeiouAEIOU"
    if contains(vowels, s[0]) {
        return 1 + CountVowels(s[1:])
    }
    return CountVowels(s[1:])
}

func contains(s string, char byte) bool {
    for i := 0; i < len(s); i++ {
        if s[i] == char {
            return true
        }
    }
    return false
}

func FactorialSum(n int) int {
    if n == 0 {
        return 0
    }
    return Factorial(n) + FactorialSum(n-1)
}