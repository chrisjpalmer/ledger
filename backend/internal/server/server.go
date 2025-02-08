package server

import "net/http"

type Server struct {
}

type Config struct {
	Addr string
	Port int
}

func NewServer(c Config) {
	srv := http.Server{}
}

func (s *Server) Close() {

}
