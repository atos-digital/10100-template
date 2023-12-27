package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

var sessionContextKey contextKey = "session"

func SessionFromContext(ctx context.Context) *sessions.Session {
	if session, ok := ctx.Value(sessionContextKey).(*sessions.Session); ok {
		return session
	}
	return nil
}

func Session(store sessions.Store, name string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), sessionContextKey, session))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
