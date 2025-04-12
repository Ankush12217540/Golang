package main

import (
    "fmt"     // Standard library for formatted I/O
    "math"    // Standard library for mathematical functions
    "math/rand" // Standard library for random number generation
)

func main() {
    // Print the value of Pi from the math package
    fmt.Println("Pi:", math.Pi)

    // Generate and print a random number between 1 and 100 using rand
    randomNumber := rand.Intn(100) + 1
    fmt.Println("Random number between 1 and 100:", randomNumber)
}


/*
✅ Key Points:
Factored imports are preferred for cleanliness and organization.
Single imports may be used when there's just one package, but it’s usually avoided in favor of the grouped approach.
Good Style: Group imports in parentheses to keep the import section neat, especially when dealing with multiple packages.
*/