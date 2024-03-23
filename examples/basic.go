package main

import (
	"fmt"

	mesquite "github.com/go-mesquite/Mesquite"
)

func main() {
	// Run a function from the other file
	message := mesquite.Hello("Partner")
	fmt.Println(message)

	m := mesquite.NewMesquite()
	m.GET()

}
