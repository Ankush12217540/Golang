package main

import "fmt"

func main() {

	// -----------------------------
	// 1️⃣ Variables (with var keyword and type)
	// -----------------------------
	// This is the traditional way to declare a variable in Go.
	var age int
	age = 25
	fmt.Println("Age:", age) // Output: Age: 25

	// -----------------------------
	// 2️⃣ Variables with Initializers
	// -----------------------------
	// You can declare a variable and assign a value at the same time.
	var city string = "Delhi"
	fmt.Println("City:", city) // Output: City: Delhi

	// -----------------------------
	// 3️⃣ Short Variable Declarations
	// -----------------------------
	// Inside a function, you can use := to declare and initialize a variable in one step.
	// This is the most commonly used form in Go for local variables.
	country := "India"
	fmt.Println("Country:", country) // Output: Country: India

	// -----------------------------
	// 4️⃣ Type Inference
	// -----------------------------
	
	// Go automatically detects the type of the value and assigns the correct type to the variable.
	// So you don't need to write the type explicitly.
	temperature := 36.5 // Go infers this as float64
	isRaining := false  // Go infers this as bool
	initial := 'A'      // Go infers this as rune (int32)

	fmt.Println("Temperature:", temperature) // Output: Temperature: 36.5
	fmt.Println("Is it raining?", isRaining) // Output: Is it raining? false
	fmt.Println("Initial:", initial)         // Output: Initial: 65 (ASCII of 'A')
}
