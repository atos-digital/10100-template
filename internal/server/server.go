package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/atos-digital/10.10.0-template/internal/config"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	r    *chi.Mux
	srv  *http.Server
	conf config.Config
}

func New(conf config.Config) (*Server, error) {
	s := new(Server)
	s.r = chi.NewRouter()
	s.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Handler: s.r,
	}
	s.conf = conf
	return s, nil
}

func (s *Server) ListenAndServe() error {
	s.middleware()
	s.Routes()
	log.Printf("server: listening on http://%s", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
