package middleware

import (
	"log"
	"net/http"
	"time"
)

// logging middleware forr detailed log of each http request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//capture the start time
		startTime := time.Now()

		//proceed to the next middleware or handler
		next.ServeHTTP(w, r)

		log.Printf("%s %s %s %v %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(startTime),
			r.Response.StatusCode)
	})
}
