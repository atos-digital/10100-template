package server

import (
	"embed"
)

//go:embed assets
var assets embed.FS

func (s *Server) Routes() {
	s.r.Handle("GET /assets/*", s.HandleAssets(assets))
	s.r.Handle("GET /favicon.ico", s.HandleFavicon(assets))

	s.r.Handle("GET /", s.handlePageIndex())

	s.srv.Handler = s.mw(s.r)
}
