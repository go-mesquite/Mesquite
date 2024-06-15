package main

import (
	"net/http"

	mesquite "github.com/go-mesquite/Mesquite"
)

/*
   This is for the controller part of the MVC
   It stores the logic for each URL endpoint
*/

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + mesquite.URLParam(r, "Message")))
}
