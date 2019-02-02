package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	Logger *log.Logger
	Srv    *http.Server
}

func (s *Server) Start() {
	err := s.Srv.ListenAndServe()
	if err != nil {
		s.Logger.Fatalf("Server is not starting due to: %v", err)
	}
	s.Logger.Printf("Server is running on port: %s", s.Srv.Addr)
}

func (s *Server) Close() {
	if err := s.Srv.Close(); err != nil {
		s.Logger.Fatalf("Error during Close: %v", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	if err := s.Srv.Shutdown(ctx); err != nil {
		s.Logger.Fatalf("Error during Shutdown: %v", err)
	}
}
