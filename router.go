package mesquite

import (
	"fmt"
	"net/http"
)

type MesquiteSite struct {
	// More configuration for the site will go here

	router Router
	// Later, add the functionality for multiple routers and names (for reverse lookups)
	// And middleware https://drstearns.github.io/tutorials/gomiddleware/
}

func NewSite() *MesquiteSite {
	site := &MesquiteSite{}
	return site
}

type Router struct {
	site *MesquiteSite
}

func NewRouter(site *MesquiteSite) *Router {
	return &Router{
		site: site,
	}
}

type Route struct {
	absolute_path string
	controller    Controller
}

// A function that can be registered to a route to handle HTTP requests
type Controller func(http.ResponseWriter, *http.Request)

func (router *Router) handleFunc(method string, path string, controller Controller) {
	// HandleFunc is the standard interface for all controllers.
}

// Serve the site
func (site *MesquiteSite) Serve(tcpAddr string) {
	fmt.Println("Serving on:", tcpAddr)
	http.ListenAndServe(tcpAddr, site.router)
}

// These HTTP method shortcuts these are the only ones HTMX supports

// Add a route and Controller for GET requests
func (router *Router) GET(path string, handle Controller) {
	//r.handle(http.MethodGet, path, handle)
	router.handleFunc("GET", path, handle)
}

// Add a route and Controller for POST requests
func (router *Router) POST(path string, handle Controller) {
	router.handleFunc("POST", path, handle)
}

// Add a route and Controller for PUT requests
func (router *Router) PUT(path string, handle Controller) {
	router.handleFunc("PUT", path, handle)
}

// Add a route and Controller for DELETE requests
func (router *Router) DELETE(path string, handle Controller) {
	router.handleFunc("DELETE", path, handle)
}

// Add a route and Controller for PATCH requests
func (router *Router) PATCH(path string, handle Controller) {
	router.handleFunc("PATCH", path, handle)
}
