package main

import "fmt"

// divideAndRemainder takes two integers and returns both the quotient and the remainder.
// The return values are *named*, which makes the code cleaner and easier to read.
func divideAndRemainder(a int, b int) (quotient int, remainder int) {
	// Calculate the quotient
	quotient = a / b

	// Calculate the remainder
	remainder = a % b

	// Since we are using named return values, we can just use 'return'
	// It will automatically return the values of 'quotient' and 'remainder'
	return
}

func main() {
	// Call the function with two numbers
	q, r := divideAndRemainder(17, 5)

	// Print the results
	fmt.Println("Quotient:", q)   // Output: Quotient: 3
	fmt.Println("Remainder:", r) // Output: Remainder: 2
}
