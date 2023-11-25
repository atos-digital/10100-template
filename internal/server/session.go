package server

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func (s *Server) cookieStore() *sessions.CookieStore {
	return &sessions.CookieStore{
		Codecs: securecookie.CodecsFromPairs(s.conf.CookieSecret),
		Options: &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
	}
}
