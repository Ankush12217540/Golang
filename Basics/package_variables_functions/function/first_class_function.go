package main

import "fmt"

// This is a normal function that takes a string and prints a greeting
func greet(name string) {
    fmt.Println("Hello,", name)
}

// This function takes another function as a parameter
// The parameter 'f' is a function that takes a string and returns nothing (void)
func callFunction(f func(string), value string) {
    f(value) // calling the function passed as argument
}

// This function returns another function
func getMultiplier(factor int) func(int) int {
    // Returning an anonymous function (closure) that multiplies by 'factor'
    return func(x int) int {
        return x * factor
    }
}

func main() {
    // Assigning a function to a variable
    myFunc := greet
    myFunc("Alice") // Output: Hello, Alice

    // Passing function as an argument
    callFunction(greet, "Bob") // Output: Hello, Bob

    // Receiving a function as a return value
    double := getMultiplier(2)
    triple := getMultiplier(3)

    fmt.Println("Double of 4:", double(4)) // Output: 8
    fmt.Println("Triple of 4:", triple(4)) // Output: 12
}

