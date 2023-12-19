package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "true" {
			http.Redirect(w, r, "/form", http.StatusMovedPermanently)
			return
		}
		var data ui.FormData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		ui.FormResult(data).Render(r.Context(), w)
	})
}

func (s *Server) handlePageSearch() http.Handler {
	return templ.Handler(ui.DefaultSearch, templ.WithContentType("text/html"))
}

func (s *Server) handleSearchUsers() http.Handler {
	data := []ui.SearchResult{
		{FirstName: "John", LastName: "Smith", Email: "johnsmith@email.com"},
		{FirstName: "Jane", LastName: "Doe", Email: "janedoe@email.com"},
		{FirstName: "Zelma", LastName: "Bush", Email: "zelmabush@email.com"},
		{FirstName: "Dorthy", LastName: "Chase", Email: "dorthychase@email.com"},
		{FirstName: "Sellers", LastName: "Carver", Email: "sellerscarver@email.com"},
		{FirstName: "Williams", LastName: "Olsen", Email: "williamsolsen@email.com"},
		{FirstName: "Florine", LastName: "Marquez", Email: "florinemarquez@email.com"},
		{FirstName: "Kathie", LastName: "Mcdowell", Email: "kathiemcdowell@email.com"},
		{FirstName: "Leach", LastName: "Alvarez", Email: "leachalvarez@email.com"},
		{FirstName: "Mitchell", LastName: "Wright", Email: "mitchellwright@email.com"},
		{FirstName: "Bridgett", LastName: "Hodge", Email: "bridgetthodge@email.com"},
		{FirstName: "Deanna", LastName: "Mcmahon", Email: "deannamcmahon@email.com"},
		{FirstName: "Freida", LastName: "Estrada", Email: "freidaestrada@email.com"},
		{FirstName: "Mamie", LastName: "Mcgee", Email: "mamiemcgee@email.com"},
		{FirstName: "Clarke", LastName: "Vazquez", Email: "clarkevazquez@email.com"},
		{FirstName: "Tanner", LastName: "Rose", Email: "tannerrose@email.com"},
		{FirstName: "Carissa", LastName: "Greene", Email: "carissagreene@email.com"},
		{FirstName: "Villarreal", LastName: "Hester", Email: "villarrealhester@email.com"},
		{FirstName: "Morrison", LastName: "Hutchinson", Email: "morrisonhutchinson@email.com"},
		{FirstName: "Collier", LastName: "Farley", Email: "collierfarley@email.com"},
		{FirstName: "Sanchez", LastName: "Dudley", Email: "sanchezdudley@email.com"},
		{FirstName: "Nichole", LastName: "Weeks", Email: "nicholeweeks@email.com"},
	}
	simpleSearch := func(query string) []ui.SearchResult {
		var results []ui.SearchResult
		for _, r := range data {
			if strings.Contains(r.FirstName, query) ||
				strings.Contains(r.LastName, query) ||
				strings.Contains(r.Email, query) {
				results = append(results, r)
			}
		}
		return results
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		query := r.FormValue("search")
		w.Header().Set("Content-Type", "text/html")
		ui.SearchResults(simpleSearch(query)).Render(r.Context(), w)
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
