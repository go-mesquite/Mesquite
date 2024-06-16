package main

import (
	"net/http"
)

/*
   This is for the controller part of the MVC
   It stores the logic for each URL endpoint
*/

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, roooooooooot"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Get a URL var
	name := r.PathValue("name")
	w.Write([]byte("Helloooooooooooooooooooo, " + name))
	//w.Write([]byte("Hello " + mesquite.URLParam(r, "Message")))
}

func aRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a route"))
}
