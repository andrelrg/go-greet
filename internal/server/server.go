package server

import (
	"github.com/andrelrg/go-greet/internal/controllers/base"
	"github.com/andrelrg/go-greet/internal/controllers/heartbeat"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Initialize() {
	s.Router = mux.NewRouter()
	s.initializeRoutes()
}

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", base.Base{heartbeat.Handler}.Serve).Methods("GET")
}
