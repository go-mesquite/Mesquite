package mesquite

import (
	"fmt"
	"net/http"
)

func Hello(name string) string {
	// Return greeting
	message := fmt.Sprintf("Howdy, %v...", name)
	return message
}

// Settings for the router
type Mesquite struct {
	setting string
}

// Handle is a function that can be registered to a route to handle HTTP
// requests. Like http.HandlerFunc, but has a third parameter for the values of
// wildcards (path variables).
type Handle func(http.ResponseWriter, *http.Request, Params)

func NewMesquite() Mesquite {
	return Mesquite{
		setting: "fish",
	}
}

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// GET is a shortcut for router.Handle(http.MethodGet, path, handle)
func (m *Mesquite) GET(path string, handle Handle) {
	m.handle(http.MethodGet, path, handle)
}

// POST is a shortcut for router.Handle(http.MethodPost, path, handle)
func (m *Mesquite) POST(path string, handle Handle) {
	m.handle(http.MethodPost, path, handle)
}

// Handles adding the route to the tree. Not made to be accessed directly
func (m *Mesquite) handle(method, path string, handle Handle) {
	fmt.Println(handle)
}

/*
NOTE: Go templates do not support HTMX. Make your own engine?
Base it off of Jinja2 to attract Django/Flask developers?
Also, integrate HTMX into the template extension (Like for )

.gohtmx or .ghtmx or .gthx or .gomx
Make this auto-format too

And maybe lose the native go handler?
See how Go Echo uses this. I'm guessing that this will end up more similar to Echo. But try to specify url type (str or int)


Ok so I've been looking at the existing Go community some more and routers are plentiful. There is no reason for me to re-invent that wheel.
So what can I build to make development faster? I can build the database and a better template layer (Checkout templ to make sure this is a need first)

*/
