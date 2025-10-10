package middleware

import (
	"net/http"
	"sync"
	"time"
)

var (
	rateLimiters = make(map[string]*rateLimiter)
	mu           sync.Mutex
)

type rateLimiter struct {
	lastRequest time.Time
	count       int
}

const (
	limit    = 10 // requests
	interval = time.Minute
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		mu.Lock()
		rl, exists := rateLimiters[ip]
		if !exists || time.Since(rl.lastRequest) > interval {
			rl = &rateLimiter{lastRequest: time.Now(), count: 1}
			rateLimiters[ip] = rl
		} else {
			rl.count++
			if rl.count > limit {
				mu.Unlock()
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			rl.lastRequest = time.Now()
		}
		mu.Unlock()
		next.ServeHTTP(w, r)
	})
}
