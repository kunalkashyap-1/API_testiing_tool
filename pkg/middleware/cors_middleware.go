package middleware

import (
	"net/http"
)

// CORSMiddleware sets the appropriate CORS headers for incoming requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow certain headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Allow methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
