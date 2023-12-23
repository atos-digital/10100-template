package server

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"

	"github.com/atos-digital/10.10.0-template/ui/pages"
)

func (s *Server) handlePageSearch() http.Handler {
	return templ.Handler(pages.DefaultSearch, templ.WithContentType("text/html"))
}

func (s *Server) handleSearchUsers() http.Handler {
	data := []pages.SearchResult{
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
	simpleSearch := func(query string) []pages.SearchResult {
		var results []pages.SearchResult
		if query == "" {
			return results
		}
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
		pages.SearchResults(simpleSearch(query)).Render(r.Context(), w)
	})
}
