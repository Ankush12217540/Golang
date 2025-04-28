package main

import "fmt"

func main() {
    // 1. Declaring a Pointer
    var p *int
    fmt.Println("Initial value of pointer p:", p) // prints <nil>

    // 2. Getting the Address of a Variable
    x := 42
    p = &x // p now holds the address of x

    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", &x)
    fmt.Println("Pointer p holds:", p)

    // 3. Dereferencing a Pointer
    fmt.Println("Value pointed by p:", *p) // Dereferencing - prints 42

    // Changing the value through pointer
    *p = 100
    fmt.Println("New value of x after changing through pointer:", x)
}
