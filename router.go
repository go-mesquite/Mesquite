package mesquite

import (
	"fmt"
	"net/http"
)

// This uses the standard library router
type Router struct {
	mux *http.ServeMux
}

// Create a new mesquite router
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// HandleFunc registers a handler function wrapper
func (router *Router) HandleFunc(httpVerb string, pattern string, handler http.HandlerFunc) {
	// Go makes "/" match with everything. This corrects the unexpected behavior
	if pattern == "/" {
		pattern = "/{$}"
	}
	// Add an http verb if there is one
	if httpVerb != "" {
		pattern = httpVerb + " " + pattern
	}
	router.mux.HandleFunc(pattern, handler)
}

// These HTTP method shortcuts these are the only ones HTMX supports

// Add a route and Controller for GET requests
func (router *Router) GET(path string, handle http.HandlerFunc) {
	//r.handle(http.MethodGet, path, handle)
	router.HandleFunc(http.MethodGet, path, handle)
}

// Add a route and Controller for POST requests
func (router *Router) POST(path string, handle http.HandlerFunc) {
	router.HandleFunc("POST", path, handle)
}

// Add a route and Controller for PUT requests
func (router *Router) PUT(path string, handle http.HandlerFunc) {
	router.HandleFunc("PUT", path, handle)
}

// Add a route and Controller for DELETE requests
func (router *Router) DELETE(path string, handle http.HandlerFunc) {
	router.HandleFunc("DELETE", path, handle)
}

// Add a route and Controller for PATCH requests
func (router *Router) PATCH(path string, handle http.HandlerFunc) {
	router.HandleFunc("PATCH", path, handle)
}

// A route that exposes a directory at a certian path. "/static" is recommended
// Note: index.html will be served at the root utlPath
func (router *Router) Static(directoryPath string, urlPath string) {
	// Create a file server for the specified directory
	fs := http.FileServer(http.Dir(directoryPath))
	fmt.Println("Register ran")

	// Strip the urlPath prefix from the request URL before passing to FileServer
	// This allows serving files relative to the given urlPath
	router.HandleFunc(http.MethodGet, urlPath+"/{filePath...}", func(w http.ResponseWriter, r *http.Request) {
		// I'm just going to count the root directory list from this as a feature..
		fmt.Println("Handle ran")
		http.StripPrefix(urlPath, fs).ServeHTTP(w, r)
	})
}

// ServeHTTP allows Router to implement the http.Handler interface so we can use the router in http.ListenAndServe()
func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	router.mux.ServeHTTP(w, req)
}
