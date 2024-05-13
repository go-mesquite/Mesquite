package mesquite

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type RouteEntry struct {
	Path        *regexp.Regexp
	Method      string
	HandlerFunc http.HandlerFunc
}

func (ent *RouteEntry) Match(r *http.Request) map[string]string {
	match := ent.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil // No match found
	}

	// Create a map to store URL parameters in
	params := make(map[string]string)
	groupNames := ent.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}

	return params
}

type Router struct {
	routes []RouteEntry
}

func NewRouter() *Router {
	return &Router{}
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	// NOTE: ^ means start of string and $ means end. Without these,
	//   we'll still match if the path has content before or after
	//   the expression (/foo/bar/baz would match the "/bar" route).
	exactPath := regexp.MustCompile("^" + path + "$")

	e := RouteEntry{
		Method:      method,
		Path:        exactPath,
		HandlerFunc: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
}

// GET shortcut
func (router *Router) GET(path string, handlerFunc http.HandlerFunc) {
	router.Route(http.MethodGet, path, handlerFunc)
}

// POST shortcut
func (router *Router) POST(path string, handlerFunc http.HandlerFunc) {
	router.Route(http.MethodPost, path, handlerFunc)
}

// PUT shortcut
func (router *Router) PUT(path string, handlerFunc http.HandlerFunc) {
	router.Route(http.MethodPut, path, handlerFunc)
}

// DELETE shortcut
func (router *Router) DELETE(path string, handlerFunc http.HandlerFunc) {
	router.Route(http.MethodDelete, path, handlerFunc)
}

// PATCH shortcut
func (router *Router) PATCH(path string, handlerFunc http.HandlerFunc) {
	router.Route(http.MethodPatch, path, handlerFunc)
}

// ServeHTTP is implemented to support the use of http.ListenAndServe()
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:", r)
			http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)
		}
	}()

	for _, e := range rtr.routes {
		params := e.Match(r)
		if params == nil {
			continue // No match found
		}

		// Create new request with params stored in context
		ctx := context.WithValue(r.Context(), "params", params)
		e.HandlerFunc.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	http.NotFound(w, r)
}

// Get the URL variable
func URLParam(r *http.Request, name string) string {
	ctx := r.Context()
	params := ctx.Value("params").(map[string]string)
	return params[name]
}

// TODO get this working

// Serves static files from the specified directory.
func (rtr *Router) Static(dir string, prefix string) {
	// Create a handler function to serve static files
	handler := http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))

	// Register the handler with the router
	rtr.Route(http.MethodGet, prefix+"/*filepath", func(w http.ResponseWriter, r *http.Request) {
		// Extract the requested file path
		filePath := URLParam(r, "filepath")

		// Validate the requested file path to prevent directory traversal
		if !strings.HasPrefix(filePath, "/") {
			http.NotFound(w, r)
			return
		}

		// Serve static files using the created handler
		handler.ServeHTTP(w, r)
	})
}
