package main

import (
    "fmt"

    "github.com/go-mesquite/Mesquite"
)

func main() {
    // Run a function from the other file
    message := mesquite.Hello("Partner")
    fmt.Println(message)
}