package internal

import (
	"context"
	"fmt"
	"net/http"
)

func NewServer(handler *http.ServeMux) Server {
	return Server{
		Port: 3000,
		handler: handler,
	}
}

type Server struct {
	Port int

	handler *http.ServeMux
	server *http.Server
}

func (s *Server) Serve() error {
	s.server = &http.Server{
		Addr: fmt.Sprintf(":%d", s.Port),
		Handler: s.handler,
	}

	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	if s.server == nil {
		return nil
	}
	return s.server.Shutdown(context.Background())
}
