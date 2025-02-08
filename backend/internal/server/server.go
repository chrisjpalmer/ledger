package server

import (
	"fmt"
	"net/http"

	openapi "github.com/chrisjpalmer/ledger/backend/internal/api/go"
)

type Server struct {
	srv *http.Server
}

type Config struct {
	Port int
}

func NewServer(c Config) {
	var s Server
	ctl := openapi.NewLedgerAPIController(&s)

	srv := http.Server{
		Handler: openapi.NewRouter(ctl),
		Addr:    fmt.Sprintf(":%d", c.Port),
	}
	s.srv = &srv
}

func (s *Server) Close() {
	s.srv.Close()
}
