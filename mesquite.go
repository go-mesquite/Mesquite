package mesquite

import "fmt"

func Hello(name string) string {
	// Return greeting
	message := fmt.Sprintf("Howdy, %v...", name)
	return message
}

/*
NOTE: Go templates do not support HTMX. Make your own engine?
Base it off of Jinja2 to attract Django/Flask developers?
Also, integrate HTMX into the template extension (Like for )

.gohtmx or .ghtmx or .gthx
Make this auto-format too

And maybe lose the native go handler?
See how Go Echo uses this. I'm guessing that this will end up more similar to Echo. But try to specify url type (str or int)
*/
