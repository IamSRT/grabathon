package middleware

import (
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
