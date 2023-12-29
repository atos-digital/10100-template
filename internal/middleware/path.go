package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var pathContextKey contextKey = "path"

func PathFromContext(ctx context.Context) string {
	if path, ok := ctx.Value(pathContextKey).(string); ok {
		return path
	}
	return ""
}

// Path adds the URL path from the request to the context
func CapturePath(next http.Handler) http.Handler {
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
