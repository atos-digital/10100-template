package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atos-digital/10.10.0-template/ui/pages"
)

func (s *Server) handlePageForm() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Header.Get("HX-Boosted") != "true" && r.Header.Get("HX-Request") == "true" {
			pages.Form().Render(r.Context(), w)
			return
		}
		pages.DefaultForm.Render(r.Context(), w)
	})
}

func (s *Server) handleFormSubmit() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "true" {
			http.Redirect(w, r, "/form", http.StatusMovedPermanently)
			return
		}
		var data pages.FormData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		pages.FormResult(data).Render(r.Context(), w)
	})
}
