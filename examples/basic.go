package main

import (
	"fmt"
	"net/http"

	mesquite "github.com/go-mesquite/Mesquite"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, root"))
}

func maion() {
	// Run a function from the other file
	message := mesquite.Hello("Partner")
	fmt.Println(message)

	// Testing the router as I build it here
	router := mesquite.NewRouter()

	router.GET("/", Root)

	// Add later
	//router.static("static". "static")
	//router.templates("templates")
	//router.404(Handle404)

	router.Serve()
}
