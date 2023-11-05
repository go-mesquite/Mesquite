# Test with jaxlo.net (import from /home/jax/Git/Mesquite)


A good place to start writing the router: https://dev.to/bmf_san/introduction-to-golang-http-router-made-with-nethttp-3nmb
Build it in the Trie data structure? (It's faster)



## Writing a router

// Go 1.22 has a built in router called mux: https://news.ycombinator.com/item?id=37898999
// Use this? Or get inspiration from it?

func add_route(path string, view fn , name string) {
    // End up with 2 maps
    // One has the URLs and the other has the names. They both then have a pointer to the view function
}


// in the user's routes.go (Use a router instead of app if possible?)
r := router.New()

r.add_route("/", routes.rootHandler, "home")
r.add_route("/about", routes.aboutHandler, "about")
r.add_route("/ok", okHandler, nil)

http.ListenAndServe("localhost:8090", r)

// Route example
def okHandler(w http.ResponseWriter, r *http.Request) {
    // Return a regular template
    return render("path/to/template")
    // Return a template and other items (HTTP2)
    return render("path/to/template").sendStatic("favicon.ico", "style.css").sendPages("root", "abut")

}



## Database
Is it a bad idea to have the dev write SQL and treat that "translation" layer as a service layer?
How do we do this while helping with initialization/migrations?
