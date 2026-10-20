package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type rateInfo struct {
	count     int
	resetTime time.Time
}

var (
	rateLimitStore = make(map[string]*rateInfo)
	rateMutex      sync.Mutex
)

const (
	maxRequestsPerMinute = 20
	windowDuration       = time.Minute
)

func RateLimit() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := clientIP(r)

			rateMutex.Lock()
			info, exists := rateLimitStore[ip]
			if !exists || time.Now().After(info.resetTime) {
				info = &rateInfo{
					count:     0,
					resetTime: time.Now().Add(windowDuration),
				}
				rateLimitStore[ip] = info
			}

			info.count++
			currentCount := info.count
			reset := info.resetTime
			rateMutex.Unlock()

			if currentCount > maxRequestsPerMinute {
				w.Header().Set("Retry-After", time.Until(reset).String())
				http.Error(w, "Muitas requisições, tente novamente mais tarde", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func clientIP(r *http.Request) string {
	// Tenta pegar IP de proxies/reverse proxies
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
