package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type contextKey string

var pathContextKey contextKey = "path"

func PathFromContext(ctx context.Context) string {
	if theme, ok := ctx.Value(pathContextKey).(string); ok {
		return theme
	}
	return ""
}

// Path adds the URL path from the request to the context
func Path(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var path string
		rctx := chi.RouteContext(r.Context())
		if rctx != nil && rctx.RoutePath != "" {
			path = rctx.RoutePath
		} else {
			path = r.URL.Path
		}
		r = r.WithContext(context.WithValue(r.Context(), pathContextKey, path))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
