package server

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/atos-digital/10.10.0-template/ui"
)

func (s *Server) handlePageData() http.Handler {
	return templ.Handler(ui.DefaultData, templ.WithContentType("text/html"))
}
