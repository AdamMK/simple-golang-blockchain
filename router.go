package main

import (
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	BlockChain []Block
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		BlockChain: []Block{},
	}
	s.requests()
	return s
}
