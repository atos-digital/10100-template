package middleware

import (
	"context"
	"net/http"
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
		r = r.WithContext(context.WithValue(r.Context(), pathContextKey, r.URL.Path))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
