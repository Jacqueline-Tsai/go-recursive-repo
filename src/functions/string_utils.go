package functions

import (
	"strings"
)

// Concatenate joins two strings together.
func Concatenate(str1, str2 string) string {
	return str1 + str2
}

// Split divides a string into a slice of substrings based on a delimiter.
func Split(str, delimiter string) []string {
	return strings.Split(str, delimiter)
}

// Reverse returns the reversed version of the input string.
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome.
func IsPalindrome(str string) bool {
	reversed := Reverse(str)
	return str == reversed
}

// CountVowels returns the number of vowels in a string.
func CountVowels(str string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

// CountConsonants returns the number of consonants in a string.
func CountConsonants(str string) int {
	consonants := "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
	count := 0
	for _, char := range str {
		if strings.ContainsRune(consonants, char) {
			count++
		}
	}
	return count
}

// ToUpper converts a string to uppercase.
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// ToLower converts a string to lowercase.
func ToLower(str string) string {
	return strings.ToLower(str)
}

// TrimSpaces removes leading and trailing spaces from a string.
func TrimSpaces(str string) string {
	return strings.TrimSpace(str)
}

// Repeat returns a string repeated n times.
func Repeat(str string, n int) string {
	return strings.Repeat(str, n)
}

// Replace replaces old substring with new substring in a string.
func Replace(str, old, new string) string {
	return strings.ReplaceAll(str, old, new)
}

// CountWords returns the number of words in a string.
func CountWords(str string) int {
	words := strings.Fields(str)
	return len(words)
}

// Recursive function to count the number of characters in a string.
func CountCharacters(str string) int {
	if str == "" {
		return 0
	}
	return 1 + CountCharacters(str[1:])
}