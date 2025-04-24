package main

import "fmt"

// Function to demonstrate a forever loop
func foreverLoop() {
    counter := 1

    // Infinite loop
    for {
        fmt.Printf("This is iteration number: %d\n", counter)
        
        // Increase the counter
        counter++

        // Stop the loop after 5 iterations (for demonstration purposes)
        if counter > 5 {
            fmt.Println("Breaking out of the loop after 5 iterations.")
            break // Exit the forever loop after 5 iterations
        }
    }
}

func main() {
    // Call the foreverLoop function to demonstrate the concept
    foreverLoop()
}
