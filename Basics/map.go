package main

import "fmt"

func main() {
    // 1. Slice from an Array
    fmt.Println("1. Slice from an Array:")
    arr := [5]int{10, 20, 30, 40, 50}
    sliceFromArray := arr[1:4] // Index 1 to 3 (20, 30, 40)
    fmt.Println("Slice from Array:", sliceFromArray)

    // Modifying slice modifies array
    sliceFromArray[0] = 99
    fmt.Println("Modified Slice:", sliceFromArray)
    fmt.Println("Modified Array:", arr)

    fmt.Println()

    // 2. Slice created directly (Literal)
    fmt.Println("2. Slice created directly (Literal):")
    directSlice := []string{"red", "green", "blue"}
    fmt.Println("Direct Slice:", directSlice)

    fmt.Println()

    // 3. Slice using make() function
    fmt.Println("3. Slice using make() function:")
    madeSlice := make([]int, 3, 5) // len=3, cap=5
    fmt.Println("Made Slice:", madeSlice)
    fmt.Println("Length:", len(madeSlice))
    fmt.Println("Capacity:", cap(madeSlice))

    fmt.Println()

    // 4. Appending elements to Slice
    fmt.Println("4. Appending elements to Slice:")
    madeSlice = append(madeSlice, 100, 200)
    fmt.Println("Slice after append:", madeSlice)
    fmt.Println("Length:", len(madeSlice))
    fmt.Println("Capacity:", cap(madeSlice)) // If capacity exceeded, Go allocates new underlying array

    fmt.Println()

    // 5. Copying one slice into another
    fmt.Println("5. Copying Slices:")
    original := []int{1, 2, 3}
    copied := make([]int, len(original))
    copy(copied, original)
    fmt.Println("Original Slice:", original)
    fmt.Println("Copied Slice:", copied)

    // Changing copied slice won't affect original
    copied[0] = 99
    fmt.Println("After modifying copied slice:")
    fmt.Println("Original Slice:", original)
    fmt.Println("Copied Slice:", copied)

    fmt.Println()

    // 6. Slicing a Slice (Sub-slicing)
    fmt.Println("6. Slicing a Slice (Sub-slicing):")
    newSlice := []int{10, 20, 30, 40, 50}
    subSlice := newSlice[2:] // from index 2 to end
    fmt.Println("Original Slice:", newSlice)
    fmt.Println("Sub Slice:", subSlice)

    fmt.Println()

    // 7. Nil Slice
    fmt.Println("7. Nil Slice:")
    var nilSlice []int
    fmt.Println("Nil Slice:", nilSlice)
    fmt.Println("Is nil?", nilSlice == nil)

    fmt.Println()

    // 8. Empty Slice (not nil, but length 0)

    
    fmt.Println("8. Empty Slice:")
    emptySlice := []int{}
    fmt.Println("Empty Slice:", emptySlice)
    fmt.Println("Is nil?", emptySlice == nil)
}
