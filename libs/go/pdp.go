package pdp

import (
	"context"
	"net/http"
	"strconv"
)

type contextKey string

const (
	// LevelKey is the context key for the PDP level.
	LevelKey contextKey = "pdp_level"
	// HeaderName is the standard HTTP header for PDP.
	HeaderName = "X-PDP-Level"
)

// Middleware parses the PDP header and injects the level into the request context.
// Invalid or missing headers fallback to the provided defaultLevel.
func Middleware(defaultLevel int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			level := defaultLevel

			if val := r.Header.Get(HeaderName); val != "" {
				if parsed, err := strconv.Atoi(val); err == nil && parsed >= 0 && parsed <= 2 {
					level = parsed
				}
			}

			ctx := context.WithValue(r.Context(), LevelKey, level)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Level extracts the parsed PDP level from the request context.
func Level(ctx context.Context) (int, bool) {
	level, ok := ctx.Value(LevelKey).(int)
	return level, ok
}
