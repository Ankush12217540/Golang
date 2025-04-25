package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
    name string
    age  int
    city string
}

func main() {
    // Creating an instance of Person
    var p1 Person
    p1.name = "Alice"
    p1.age = 25
    p1.city = "New York"

    // Another way using struct literal
    p2 := Person{name: "Bob", age: 30, city: "Los Angeles"}

    // Print values
    fmt.Println("Person 1:", p1)
    fmt.Println("Person 2:", p2)

    // Access individual fields
    fmt.Println(p1.name, "is", p1.age, "years old and lives in", p1.city)
}
