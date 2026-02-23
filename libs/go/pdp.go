package pdp

import (
	"context"
	"net/http"
	"strconv"
)

// Level represents the data privacy classification (0-2).
type Level int

const (
	LevelPrivate  Level = 0
	LevelPersonal Level = 1
	LevelGlobal   Level = 2
)

// HeaderName is the standard HTTP header key.
const HeaderName = "X-PDP-Level"

type ctxKey struct{}

func (l Level) String() string {
	switch l {
	case LevelPrivate:
		return "Private"
	case LevelPersonal:
		return "Personal"
	case LevelGlobal:
		return "Global"
	default:
		return "Unknown"
	}
}

// Middleware injects the PDP level into the request context.
// It defaults to LevelPrivate if the header is missing or invalid.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level := LevelPrivate

		if val := r.Header.Get(HeaderName); val != "" {
			if i, err := strconv.Atoi(val); err == nil {
				if l := Level(i); l >= LevelPrivate && l <= LevelGlobal {
					level = l
				}
			}
		}

		ctx := context.WithValue(r.Context(), ctxKey{}, level)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FromContext returns the PDP level from a context, defaulting to LevelPrivate.
func FromContext(ctx context.Context) Level {
	if level, ok := ctx.Value(ctxKey{}).(Level); ok {
		return level
	}
	return LevelPrivate
}
