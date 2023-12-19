package server

import (
	"encoding/json"
	"fmt"
	"log"
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
	return templ.Handler(ui.DefaultHome, templ.WithContentType("text/html"))
}

func (s *Server) handlePageForm() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Header.Get("HX-Boosted") != "true" && r.Header.Get("HX-Request") == "true" {
			ui.Form().Render(r.Context(), w)
			return
		}
		ui.DefaultForm.Render(r.Context(), w)
	})
}

func (s *Server) handleFormSubmit() http.Handler {
	type form struct {
		Message string `json:"message"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "true" {
			http.Redirect(w, r, "/form", http.StatusMovedPermanently)
			return
		}
		var f form
		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		ui.FormResult(f.Message).Render(r.Context(), w)
	})
}

func (s *Server) handlePageContact() http.Handler {
	return templ.Handler(ui.DefaultContact, templ.WithContentType("text/html"))
}

func (s *Server) handlePageAbout() http.Handler {
	return templ.Handler(ui.DefaultAbout, templ.WithContentType("text/html"))
}

func (s *Server) handleSaveSession() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sess.Get(r, s.conf.CookieName)
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
