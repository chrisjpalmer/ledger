package server

import (
	"fmt"
	"net/http"

	openapi "github.com/chrisjpalmer/ledger/backend/internal/api/go"
	"github.com/chrisjpalmer/ledger/backend/internal/postgres"
	"go.uber.org/zap"
)

type Server struct {
	httpSrv *http.Server
	port    int
	pgs     *postgres.Postgres
	zl      *zap.Logger
}

type Config struct {
	Port int
}

func NewServer(zl *zap.Logger, postgres *postgres.Postgres, c Config) *Server {
	srv := Server{
		port: c.Port,
		pgs:  postgres,
		zl:   zl,
	}

	// configure server
	ctl := openapi.NewLedgerAPIController(&srv)
	srv.httpSrv = &http.Server{
		Handler: openapi.NewRouter(ctl),
		Addr:    fmt.Sprintf(":%d", c.Port),
	}

	return &srv
}

func (s *Server) Listen() error {
	// start listening
	s.zl.Info("server listening...", zap.Int("port", s.port))
	err := s.httpSrv.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}

	return nil
}

func (s *Server) Close() error {
	s.zl.Info("server closing...")
	return s.httpSrv.Close()
}
