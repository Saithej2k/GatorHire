package middleware

import (
	"net/http"
)

// Apply applies all middleware to a handler
func Apply(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}