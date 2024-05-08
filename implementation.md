# Notes on the implementation

## General
- Regex is a useful tool for URLs, input validation (And SQL?). This framework fully embraces Regex to keep with the theme of simplicity at the expense of ease
- How should a website be laid out? https://go.dev/doc/modules/layout


## Router
- Made to be interchangeable with any standard Go router. Based on https://codesalad.dev/blog/how-to-build-a-go-router-from-scratch-3
- It uses a map to store routes. Optimize later with a prefix tree.
- Only GET, POST, PUT, DELETE, PATCH are supported because that is what HTMX supports. If other methods are needed, those can be added later
- Eventually I would like to add the ability to have multiple routers that can be nested. But this proved to be too much of a time sink so I'll come back to it later if there is a need.


## Middleware
- TODO add middleware support
- Add middleware that shows request speeds
- Add middleware to remove trailing slashes in URLs


## Database
Is it a bad idea to have the dev write SQL and treat that "translation" layer as a service layer?
How do we do this while helping with initialization/migrations?





## Wishlist (My brain needs to put this somewhere)
- Only support HTTPS? (Setup a way to test in dev and take advantage of HTTP2 features?)
- Add an easy way to have server side analytics. Middleware after the request is sent?
- Figure out how to have a simple que. And timed events?
- Use ko.build?
