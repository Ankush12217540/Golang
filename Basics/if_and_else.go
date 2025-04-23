package main

import "fmt"

func main() {
	// ✅ Basic If statement
	fmt.Println("1. Basic If Statement:")
	age := 20
	if age >= 18 {
		fmt.Println("  You are eligible to vote.")
	}

	// ✅ If with a short statement
	// This is useful when you want to declare a variable for just the if block
	fmt.Println("\n2. If with Short Statement:")
	if length := len("Golang"); length > 5 {
		fmt.Printf("  The word 'Golang' has %d characters.\n", length)
	}

	// ✅ If and else block
	fmt.Println("\n3. If and Else Statement:")
	number := -10
	if number >= 0 {
		fmt.Println("  The number is non-negative.")
	} else {
		fmt.Println("  The number is negative.")
	}

	// ✅ If-else ladder for multiple conditions
	fmt.Println("\n4. If-Else Ladder Example:")
	score := 75
	if score >= 90 {
		fmt.Println("  Grade: A")
	} else if score >= 80 {
		fmt.Println("  Grade: B")
	} else if score >= 70 {
		fmt.Println("  Grade: C")
	} else {
		fmt.Println("  Grade: D or below")
	}
}
