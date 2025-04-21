package main

import "fmt"

func main() {
    fmt.Println("Start")

    // This line will be deferred until the end of main()
    defer fmt.Println("Deferred: This prints last")

    fmt.Println("End")
}
