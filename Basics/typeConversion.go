package main

import "fmt"

func main() {

	// --------------------------------
	// 1️⃣ Type Conversion (Explicit)
	// --------------------------------
	// Go does NOT allow implicit type conversion between different types.
	// You must explicitly convert one type to another.

	var a int = 10
	var b float64 = 4.5

	// ❌ Invalid: cannot do a + b directly (int + float64)
	// var result = a + b // This will cause a compile-time error

	// ✅ Convert 'a' to float64 before adding
	var result = float64(a) + b
	fmt.Println("Result of float64(a) + b =", result) // Output: 14.5

	// --------------------------------
	// 2️⃣ Integer Division & Conversion
	// --------------------------------
	var x int = 7
	var y int = 2

	var intDivision = x / y // int division (result: 3, not 3.5)
	fmt.Println("Integer division x / y =", intDivision)

	// To get a float result, convert one operand to float
	var floatDivision = float64(x) / float64(y)
	fmt.Println("Float division float64(x)/float64(y) =", floatDivision)

	// --------------------------------
	// 3️⃣ Converting Between Basic Types
	// --------------------------------
	var i int = 65
	var r rune = rune(i)       // int to rune
	var s string = string(r)   // rune to string (char representation)

	fmt.Println("Integer i:", i)
	fmt.Println("Converted to rune:", r)
	fmt.Println("Rune to string (char):", s) // Output: "A"

	// --------------------------------
	// 4️⃣ Byte and String Conversions
	// --------------------------------
	str := "GoLang"
	bytes := []byte(str) // convert string to byte slice
	fmt.Println("String to byte slice:", bytes)

	strBack := string(bytes)
