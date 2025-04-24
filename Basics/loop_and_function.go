package main

import "fmt"

// Function that calculates the sum of two numbers
func add(a int, b int) int {
    return a + b
}

// Function that prints numbers from 1 to a given limit using a for loop
func printNumbers(limit int) {
    for i := 1; i <= limit; i++ {
        fmt.Println(i)
    }
}

// Function that demonstrates a while-like loop (using for loop)
func printEvenNumbers(limit int) {
    i := 2
    for i <= limit {
        fmt.Println(i)
        i += 2
    }
}

// Function that demonstrates a loop with break and continue
func demonstrateBreakContinue(limit int) {
    for i := 1; i <= limit; i++ {
        if i == 5 {
            fmt.Println("Skipping 5")
            continue // Skip the rest of the loop when i is 5
        }
        if i == 8 {
            fmt.Println("Stopping the loop at 8")
            break // Exit the loop when i is 8
        }
        fmt.Println(i)
    }
}

func main() {
    // Demonstrating the add function
    result := add(10, 5)
    fmt.Println("Addition Result:", result)

    // Demonstrating loop using printNumbers function
    fmt.Println("Printing numbers from 1 to 5:")
    printNumbers(5)

    // Demonstrating loop for even numbers using printEvenNumbers function
    fmt.Println("Printing even numbers up to 10:")
    printEvenNumbers(10)

    // Demonstrating break and continue inside a loop
    fmt.Println("Demonstrating break and continue:")
    demonstrateBreakContinue(10)
}
