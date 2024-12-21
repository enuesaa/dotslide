package internal

import (
	"context"
	"fmt"
	"net/http"
)

func NewServer(handler http.Handler) Server {
	return Server{
		Port: 3000,
		handler: handler,
	}
}

type Server struct {
	Port int

	handler http.Handler
	server *http.Server
}

func (s *Server) Serve() error {
	s.server = &http.Server{
		Addr: fmt.Sprintf(":%d", s.Port),
		Handler: s.handler,
	}
	fmt.Println("server started")

	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	if s.server == nil {
		return nil
	}
	return s.server.Shutdown(context.Background())
}
