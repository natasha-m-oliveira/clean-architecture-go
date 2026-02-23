package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(router http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}

func (s *Server) WithPort(port int) *Server {
	s.httpServer.Addr = fmt.Sprintf(":%d", port)
	return s
}

func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Server started on port", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	fmt.Println("Server gracefully stopped")
	return nil
}
