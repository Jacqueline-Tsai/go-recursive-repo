// filepath: /code_craft/go-recursive-repo/src/functions/math_utils.go
package functions

// Add returns the sum of two integers.
func Add(a int, b int) int {
    return a + b
}

// Subtract returns the difference of two integers.
func Subtract(a int, b int) int {
    return a - b
}

// Multiply returns the product of two integers.
func Multiply(a int, b int) int {
    return a * b
}

// Divide returns the quotient of two integers. It returns 0 if the divisor is 0.
func Divide(a int, b int) int {
    if b == 0 {
        return 0 // Avoid division by zero
    }
    return a / b
}

// Power returns the result of raising base to the exponent.
func Power(base int, exponent int) int {
    if exponent == 0 {
        return 1
    }
    return base * Power(base, exponent-1)
}

// Modulus returns the remainder of the division of two integers.
func Modulus(a int, b int) int {
    if b == 0 {
        return 0 // Avoid division by zero
    }
    return a % b
}

// Average returns the average of a slice of integers.
func Average(numbers []int) float64 {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return float64(sum) / float64(len(numbers))
}

// Factorial returns the factorial of a non-negative integer n.
func Factorial(n int) int {
    if n < 0 {
        return -1 // Factorial is not defined for negative numbers
    }
    if n == 0 {
        return 1
    }
    return n * Factorial(n-1)
}

// GCD returns the greatest common divisor of two integers using recursion.
func GCD(a int, b int) int {
    if b == 0 {
        return a
    }
    return GCD(b, a%b)
}

// LCM returns the least common multiple of two integers.
func LCM(a int, b int) int {
    return (a * b) / GCD(a, b)
}

// IsPrime checks if a number is prime.
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

// Square returns the square of an integer.
func Square(n int) int {
    return n * n
}

// Cube returns the cube of an integer.
func Cube(n int) int {
    return n * n * n
}

// Sum returns the sum of a slice of integers.
func Sum(numbers []int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}