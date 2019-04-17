package auth

import (
	"context"
	"net/http"
)

func Middleware() func(http.Handler) http.Handler {
	var Authorization string
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), Authorization, r.Header.Get("Authorization"))
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
