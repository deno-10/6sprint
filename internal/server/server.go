package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

const addr = ":8080"

type Server struct {
	logger *log.Logger
	srv    *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("POST /upload", handlers.Upload)

	return &Server{
		logger: logger,
		srv: &http.Server{
			Addr:         addr,
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	s.logger.Printf("сервер слушает http://localhost%s", addr)

	return s.srv.ListenAndServe()
}
