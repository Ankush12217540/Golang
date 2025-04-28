package main

import "fmt"

func main() {
    // 1. Declaring a Pointer
    var p *int // p is a pointer to an int
    fmt.Println("Initial value of pointer p:", p) // will print: <nil>

    // 2. Getting the Address of a Variable
    x := 42
    p = &x // p now holds the address of x

    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", &x)
    fmt.Println("Pointer p holds:", p)
}
