package middleware

import (
	"net/http"

	"github.com/pithandev/reverse-proxy/internal/ratelimit"
)

func RateLimit() Middleware {

	limiter := ratelimit.NewLimiter()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.RemoteAddr

			if !limiter.Allow(key) {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("rate limit exceeded!"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
