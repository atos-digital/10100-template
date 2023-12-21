package server

import (
	"net/http"

	"github.com/atos-digital/10.10.0-template/ui"
)

func (s *Server) handlePageData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		ui.DefaultData.Render(r.Context(), w)
	})
}
