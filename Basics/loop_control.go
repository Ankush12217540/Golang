package main

import "fmt"

func main() {
	// ✅ For – Standard loop with init, condition, post
	fmt.Println("Standard For Loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  Count: %d\n", i)
	}

	// ✅ For continued – use of `continue` in loop
	fmt.Println("\nFor Loop with Continue (Skip Even Numbers):")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("  Odd Number: %d\n", i)
	}

	// ✅ For is Go's "while" – act like a while loop
	fmt.Println("\nFor as While Loop (Countdown):")
	x := 5
	for x > 0 {
		fmt.Printf("  Countdown: %d\n", x)
		x--
	}

	// ✅ Forever – Infinite loop using `for {}` (demo with break)
	fmt.Println("\nInfinite Loop with Break:")
	y := 1
	for {
		fmt.Printf("  Infinite Count: %d\n", y)
		if y >= 3 {
			break // manually break the infinite loop
		}
		y++
	}

	// ✅ Exercise: Loops and Functions – sum of squares using loop + function
	fmt.Println("\nExercise: Sum of Squares Function using Loop")
	sum := sumOfSquares(5) // 1² + 2² + 3² + 4² + 5² = 55
	fmt.Printf("  Sum of Squares from 1 to 5: %d\n", sum)
}

// Function to calculate sum of squares up to n
func sumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i * i
	}
	return total
}
