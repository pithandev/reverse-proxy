package server

import "net/http"

type Server struct {
	addr    string
	handler http.Handler
}

func New(addr string, handler http.Handler) *Server {
	return &Server{
		addr:    addr,
		handler: handler,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.handler)
}
