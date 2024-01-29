package server

import (
	"net/http"
	"rest_project/internal/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	storage *storage.Storage
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	s.configureRouter()
	if err := s.configureStorage(); err != nil {
		return err
	}
	
	return http.ListenAndServe(":4060", s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/person", s.personGetAll()).Methods("GET")
	s.router.HandleFunc("/person/{id}", s.personGet()).Methods("GET")
	s.router.HandleFunc("/person", s.personSet()).Methods("POST")
	s.router.HandleFunc("/person/{id}", s.personUpdate()).Methods("PUT")
	s.router.HandleFunc("/person/{id}", s.personDelete()).Methods("DELETE")
}

func (s *Server) configureStorage() error {
	store := storage.NewStorage()
	if err := store.Open(); err != nil {
		return err
	}

	s.storage = store
	return nil
}
