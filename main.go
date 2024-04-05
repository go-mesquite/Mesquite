package mesquite

import (
	"fmt"
)

func Hello(name string) string {
	// An example of running a function
	message := fmt.Sprintf("Howdy, %v...", name)
	return message
}

/*
NOTE: Go templates do not support HTMX. Make your own template engine?
Base it off of Jinja2 to attract Django/Flask developers?
Also, integrate HTMX into the template extension (Like for )

.gohtmx or .ghtmx or .gthx or .gomx
Make this auto-format too

And maybe lose the native go handler?
See how Go Echo uses this. I'm guessing that this will end up more similar to Echo. But try to specify url type (str or int)


Ok so I've been looking at the existing Go community some more and routers are plentiful. There is no reason for me to re-invent that wheel.
So what can I build to make development faster? I can build the database and a better template layer (Checkout templ to make sure this is a need first)

*/
