package mesquite

import "net/http"

type Middleware func(http.Handler) http.Handler

// Middleware chaining. From: https://youtu.be/H7tbjKFSg58?feature=shared&t=586
func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}

}

// TODO add this as a handler
