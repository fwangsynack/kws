package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func WithValue(key, val interface{}) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), key, val))
			fmt.Printf("middleware WithValue called - %T\n", next)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
