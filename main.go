package mesquite

import "net/http"

/*
NOTE: Go templates do not support HTMX. Make your own template engine? Or use templ
Base it off of Jinja2 to attract Django/Flask developers?
Also, integrate HTMX into the template extension

.gohtmx or .ghtmx or .gthx or .gomx
Make this auto-format too

And maybe lose the native go handler?
See how Go Echo uses this. I'm guessing that this will end up more similar to Echo. But try to specify url type (str or int)


Ok so I've been looking at the existing Go community some more and routers are plentiful. There is no reason for me to re-invent that wheel.
So what can I build to make development faster? I can build the database and a better template layer (Checkout templ to make sure this is a need first)

*/

// Get the URL variable
func URLParam(r *http.Request, name string) string {
	ctx := r.Context()
	params := ctx.Value("params").(map[string]string)
	return params[name]
}
