package main

import "github.com/chrisjpalmer/ledger/backend/internal/server"

func main() {
	srv := server.NewServer(cfg)

	srv.Close()
}
