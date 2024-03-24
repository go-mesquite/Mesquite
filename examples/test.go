package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// Handler for the root path ("/")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Fprintf(w, "This is the index page!")
		} else {
			http.NotFound(w, r)
		}
	})

	// Handler for all other paths
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a different page for path: %s", r.URL.Path)
	})

	// Serve the app
	fmt.Println("Serving on localhost:8000")
	http.ListenAndServe("localhost:8000", router)
}
