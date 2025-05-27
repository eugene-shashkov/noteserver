package noteserver

import "github.com/gorilla/mux"

type NoteServer interface {
	CreateServer(sv *Server) NoteServer
	Routing() *mux.Router
}

type Server struct{}

func (s Server) CreateServer(sv *Server) NoteServer {
	return sv
}
