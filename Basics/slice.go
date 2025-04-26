package main

import "fmt"

func main() {
    // Basic array
    arr := [5]int{10, 20, 30, 40, 50}

    // Creating a slice from the array
    // Syntax: array[startIndex:endIndex] => includes startIndex, excludes endIndex
    slice1 := arr[1:4] // elements at index 1, 2, and 3 (20, 30, 40)
    fmt.Println("slice1:", slice1)

    // Slices are like references to arrays
    // Modifying slice will modify the underlying array too
    slice1[0] = 99 // Changing the first element of slice1
    fmt.Println("After modifying slice1:")
    fmt.Println("slice1:", slice1)
    fmt.Println("arr:", arr) // arr is also modified

    // Creating a slice directly (without using an array)
    slice2 := []string{"apple", "banana", "cherry"}
    fmt.Println("slice2:", slice2)

    // Appending elements to slice
    slice2 = append(slice2, "date", "elderberry")
    fmt.Println("slice2 after append:", slice2)

    // Length and capacity of slice
    fmt.Println("Length of slice2:", len(slice2))
    fmt.Println("Capacity of slice2:", cap(slice2))

    // Slicing a slice
    slice3 := slice2[1:4] // banana, cherry, date
    fmt.Println("slice3 (sliced from slice2):", slice3)

    // Make function to create a slice
    slice4 := make([]int, 5) // slice with length 5 and default values (0)
    fmt.Println("slice4:", slice4)

    // Changing elements
    slice4[0] = 100
    slice4[1] = 200
    fmt.Println("slice4 after modification:", slice4)

    // Copying slices
    slice5 := make([]int, len(slice4))
    copy(slice5, slice4)
    fmt.Println("slice5 (copy of slice4):", slice5)

    // Modifying slice5 won't affect slice4
    slice5[0] = 500
    fmt.Println("After modifying slice5:")
    fmt.Println("slice5:", slice5)
    fmt.Println("slice4:", slice4)
}
