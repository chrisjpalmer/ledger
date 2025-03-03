package server

import (
	"fmt"
	"net/http"
	"time"

	openapi "github.com/chrisjpalmer/ledger/backend/internal/api/go"
	"github.com/chrisjpalmer/ledger/backend/internal/postgres"
	"github.com/gorilla/mux"
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
		Handler: newRouter(zl, ctl),
		Addr:    fmt.Sprintf(":%d", c.Port),
	}

	return &srv
}

// newRouter - wire the openapi generated code to gorilla mux
func newRouter(zl *zap.Logger, api openapi.Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for name, route := range api.Routes() {
		var handler http.Handler = route.HandlerFunc
		handler = logger(zl, handler, name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(name).
			Handler(handler)
	}

	return router
}

func logger(zl *zap.Logger, inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		zl.Info(
			"request handled",
			zap.String("method", r.Method),
			zap.String("request_uri", r.RequestURI),
			zap.String("name", name),
			zap.Duration("duration", time.Since(start)),
		)
	})
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
