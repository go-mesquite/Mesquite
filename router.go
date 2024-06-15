package mesquite

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
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

// Serve a single static file
func (router *Router) StaticFile(filePath string, urlPath string) {
	router.Route(http.MethodGet, urlPath, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})
}

// TODO come back and figure this out after setting up variables to replace regex
// Serves static files from the specified directory.
func (router *Router) StaticDirectory(directoryPath string, urlPath string) {
	// Create a file server for the specified directory
	fs := http.FileServer(http.Dir(directoryPath))

	// Strip the urlPath prefix from the request URL before passing to FileServer
	// This allows serving files relative to the given urlPath
	router.Route(http.MethodGet, urlPath+`/{str:filePath}`, func(w http.ResponseWriter, r *http.Request) {
		// Extract the filePath from the URL path parameter
		filePath := URLParam(r, "filePath")
		fmt.Println(filePath)

		// Combine the directory and file path to get the full file path
		//fullPath := path.Join(directoryPath, filePath)

		// Check if the file exists and serve it using the file server
		// If the file doesn't exist, FileServer will handle returning a 404
		http.StripPrefix(urlPath, fs)
		fs.ServeHTTP(w, r)
	})
}

/*
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

func (router *Router) Static(path string, handlerFunc http.HandlerFunc) {
	router.Route(http.MethodGet, path, func(
		// Handle staticfile match
	))
}

func (router *Router) Static(prefix, directory string) {
	// Ensure prefix ends with a slash
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	// Register a route for serving static files
	router.Route("GET", prefix+"*", func(
		// StaticFileHandler returns an http.HandlerFunc that serves static files
		// from the specified directory.
		dir := http.Dir(directory)
		fs := http.FileServer(dir)

		return func(w http.ResponseWriter, r *http.Request) {
			// Get the file path from the URL
			filePath := path.Join(directory, r.URL.Path)

			// Check if the file exists
			_, err := os.Stat(filePath)
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}

			// Serve the file using the standard file server
			fs.ServeHTTP(w, r)
		}
	))
}




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
*/
