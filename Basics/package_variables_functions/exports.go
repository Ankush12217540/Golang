// Code Example to Demonstrate Exported and Unexported Names
package main

import (
	"fmt"
	"math" // Importing the math package
)

func main() {
	// Accessing an exported name from the math package
	fmt.Println("Value of Pi:", math.Pi) // Correct, Pi is exported from the math package

	// Attempting to access an unexported name from the math package
	// This will cause an error, as 'pi' is unexported
	// Uncomment the following line to see the error:
	// fmt.Println("Value of Pi:", math.pi) // Error: 'pi' is unexported in math package
}

/*
Summary

Exported Names:
Start with an uppercase letter.
Accessible from other packages.
Example: MyVariable, MyFunction.

Unexported Names:
Start with a lowercase letter.
Not accessible from other packages.
Example: myVariable, myFunction.

Imported Names:
These are the names you access from external packages after importing them using the import keyword.
You can only access the exported names of the imported package (those starting with an uppercase letter).
Example: fmt.Println(), math.Pi.


*/
