package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/atos-digital/10.10.0-template/ui"
)

func (s *Server) HandleFavicon() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := assets.ReadFile("assets/img/favicon.ico")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(b)
	})
}

func (s *Server) handlePageIndex() http.Handler {
	name := "World!"
	return templ.Handler(ui.Index(name), templ.WithContentType("text/html"))
}

func (s *Server) handleSaveSession() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sess.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["foo"] = "bar"
		session.Values[42] = 43
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func (s *Server) handleReadSession() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sess.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, session.Values)
	})
}
