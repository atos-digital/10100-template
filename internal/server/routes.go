package server

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed assets
var assets embed.FS

func (s *Server) Routes() error {
	// special case: handler assets content for Chi router with subroute, default go router in 1.22 will not require this step
	contentAssets, err := fs.Sub(fs.FS(assets), "assets")
	if err != nil {
		return err
	}
	s.r.Method(http.MethodGet, "/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(contentAssets))))
	s.r.Method(http.MethodGet, "/favicon.ico", s.HandleFavicon())

	s.r.Method(http.MethodGet, "/session", s.handleSaveSession())
	s.r.Method(http.MethodGet, "/read-session", s.handleReadSession())

	s.r.Method(http.MethodGet, "/", s.handlePageIndex())

	s.r.Method(http.MethodGet, "/form", s.handlePageForm())
	s.r.Method(http.MethodGet, "/form/submit", s.handleFormSubmit())
	s.r.Method(http.MethodPost, "/form/submit", s.handleFormSubmit())

	s.r.Method(http.MethodGet, "/search", s.handlePageSearch())
	s.r.Method(http.MethodPost, "/search/users", s.handleSearchUsers())

	s.r.Method(http.MethodGet, "/data", s.handlePageData())

	s.r.Method(http.MethodGet, "/contact", s.handlePageContact())
	s.r.Method(http.MethodGet, "/about", s.handlePageAbout())

	return nil
}
