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
	s.r.Method(http.MethodGet, "/", s.handlePageIndex())

	return nil
}
