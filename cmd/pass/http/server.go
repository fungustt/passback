package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type (
	Server struct {
		server *http.Server
	}

	Handler interface {
		Route() string
		ServeHTTP(http.ResponseWriter, *http.Request)
		Methods() string
	}
)

func NewServer(addr string, handlers []Handler) *Server {
	router := mux.NewRouter()
	for _, handler := range handlers {
		router.HandleFunc(handler.Route(), handler.ServeHTTP).Methods(handler.Methods())
	}

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *Server) Start(ctx context.Context) {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen and server error: %s\n", err)
		}
	}()

	log.Printf("server listen started at %s", s.server.Addr)

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server shutdown failed: %s", err)
	}
}
