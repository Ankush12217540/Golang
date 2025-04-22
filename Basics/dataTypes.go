package main

import "fmt"

func main() {

	// --------------------------------
	// 1️⃣ Basic Types in Go
	// --------------------------------
	// These are the most common types used in Go.
	var name string = "Alice"  // string
	var age int = 30           // integer
	var height float64 = 5.8   // floating-point number
	var isStudent bool = false // boolean
	var grade rune = 'A'       // rune (Unicode character, internally int32)
	var smallNum byte = 255    // byte is alias for uint8

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Grade:", grade) // Will print ASCII value of 'A' (i.e., 65)
	fmt.Println("Small Number:", smallNum)

	// --------------------------------
	// 2️⃣ Zero Values
	// --------------------------------
	// Variables that are declared but not initialized are assigned a "zero value"
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Println("\nZero value of int:", defaultInt)     // 0
	fmt.Println("Zero value of float64:", defaultFloat) // 0
	fmt.Println("Zero value of bool:", defaultBool)     // false
	fmt.Println("Zero value of string:", defaultString) // empty string ""

	// --------------------------------
	// 3️⃣ Constants
	// --------------------------------
	// Constants are declared using the 'const' keyword and cannot be changed later.
	const pi = 3.14159
	const language = "GoLang"

	// pi = 3.14 // ❌ Not allowed: constants cannot be reassigned
	fmt.Println("\nConstant pi:", pi)
	fmt.Println("Constant language:", language)

	// --------------------------------
	// 4️⃣ Numeric Constants
	// --------------------------------
	// Numeric constants are untyped until used.
	const x = 100     // x is a numeric constant
	const y = 3.14    // y is also numeric and can fit into many types
	const z = x + 1.5 // still valid even mixing types because they're constants

	// You can assign to any compatible numeric type without error
	var floatVal float64 = y
	var intVal int = x

	fmt.Println("\nNumeric Constant x:", x)
	fmt.Println("Numeric Constant y:", y)
	fmt.Println("Numeric Constant x + 1.5 =", z)
	fmt.Println("Assigned to float64:", floatVal)
	fmt.Println("Assigned to int:", intVal)
}
