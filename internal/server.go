package internal

import (
	"context"
	"fmt"
	"net/http"
)

func NewServer(router Router) Server {
	return Server{
		Port: 3000,
		router: router,
	}
}

type Server struct {
	Port int

	router Router
	server *http.Server
}

func (s *Server) Serve() error {
	s.server = &http.Server{
		Addr: fmt.Sprintf(":%d", s.Port),
		Handler: s.router.Handle(),
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
