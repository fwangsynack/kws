package middleware

import (
	"fmt"
	//"log"
	"net/http"
	//"time"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//start := time.Now()
		//log.Printf("Started %s %s", r.Method, r.URL.Path)
		fmt.Printf("middleware Logger called - %T\n", next)
		next.ServeHTTP(w, r)
		//log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}
	return http.HandlerFunc(fn)
}
