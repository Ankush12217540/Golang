package main // The entry point of the program is in the 'main' package

import (
	"fmt"         // Import the fmt package for formatted I/O
	"math/rand"   // Import the rand package for generating random numbers
	"time"        // Import the time package to work with time-related functions
)

// main function - Entry point of the Go program
func main() {
	// Using the 'fmt' package to print output
	fmt.Println("Welcome to Go Packages Tutorial!")

	// Using the 'math/rand' package to generate a random number
	// Seed the random number generator with current time to get different random numbers each time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	randomNumber := rand.Intn(100) + 1
	fmt.Println("A random number between 1 and 100:", randomNumber)

	// Calling another function that uses the 'fmt' package
	printGreeting("Go Developer")

	// Calling a function that generates a random number
	randomColor()
}

// Example function that uses the 'fmt' package to print a greeting
func printGreeting(name string) {
	// Printing using fmt package
	fmt.Println("Hello,", name)
}

// Example function to print a random color using the 'rand' package
func randomColor() {
	colors := []string{"Red", "Green", "Blue", "Yellow", "Purple"}
	// Generate a random index to pick a color
	randomIndex := rand.Intn(len(colors))
	fmt.Println("Random Color:", colors[randomIndex])
}
