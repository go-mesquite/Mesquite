package main

import (
	"fmt"
	"net/http"

	mesquite "github.com/go-mesquite/Mesquite"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, root"))
}

func Ok(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, ok"))
}

/*
// --- ---- --- Middleware --- --- ---
// Logger is a middleware handler that does request logging

	type Logger struct {
		handler http.Handler
	}

// ServeHTTP handles the request by passing it to the real handler and logging the request details

	func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		// Start the timer to see how long the controller took to execute
		startTime := time.Now()

		// Run the handler/controller
		l.handler.ServeHTTP(w, r)

		// Stop the timer and get how long the controller took to run
		duration := time.Since(startTime)
		milliseconds := duration.Milliseconds()

		// Instead of printing "0ms", we should print <1ms if it executed quickly
		if milliseconds == 0 {
			log.Printf("\x1b[32m%s %s <%dms\x1b[0m\n", r.Method, r.URL.Path, 1)
		} else {
			log.Printf("\x1b[32m%s %s %dms\x1b[0m\n", r.Method, r.URL.Path, milliseconds)
		}
	}

// NewLogger constructs a new Logger middleware handler

	func NewLogger(handlerToWrap http.Handler) *Logger {
		return &Logger{handlerToWrap}
	}
*/
func main() {
	// Run a function from the other file
	message := mesquite.Hello("Partner")
	fmt.Println(message)

	// Go Mesquite!
	site := mesquite.NewSite()
	router := mesquite.NewRouter(site)

	// TODO make my own router. The standard lib one is annoying
	// Have Routes, Methods, and Names. Could it be as easy as having a hash table for each of them?

	router.GET("/ok", Ok)
	router.GET("/", Root)

	// Add later
	//router.UseMiddleware(Func)
	//router.Static("static". "static")
	//router.Templates("templates")
	//router.404(Handle404)

	router.Serve()
}
