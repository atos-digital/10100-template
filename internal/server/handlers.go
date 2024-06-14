package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/atos-digital/10100-template/ui/pages"
)

func (s *Server) HandleAssets(assets embed.FS) http.Handler {
	contentAssets, err := fs.Sub(fs.FS(assets), "assets")
	if err != nil {
		log.Fatalf("HandleAssets: failed to load assets: %v", err)
	}
	return http.StripPrefix("/assets/", http.FileServerFS(contentAssets))
}

func (s *Server) HandleFavicon(assets embed.FS) http.Handler {
	b, err := assets.ReadFile("assets/img/favicon.ico")
	if err != nil {
		log.Fatalf("HandleFavicon: failed to read favicon.ico: %v", err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		_, err = w.Write(b)
		if err != nil {
			log.Printf("HandleFavicon: failed to write favicon.ico: %v", err)
		}
	})
}

func (s *Server) handlePageIndex() http.Handler {
	return templ.Handler(pages.DefaultHome, templ.WithContentType("text/html"))
}
