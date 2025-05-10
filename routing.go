package main

import (
	"github.com/eugene-shashkov/notema-server/handlers"

	"github.com/gorilla/mux"
)

// Routing func is a startpoint of routing mechanism
func Routing() *mux.Router {
	mx := mux.NewRouter()
	mx.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	mx.HandleFunc("/registration", handlers.RegistrationHandler).Methods("POST").Queries("email", "{email}", "password", "{password}", "name", "{name}")
	mx.HandleFunc("/login", handlers.LoginHandler).Methods("POST").Queries("email", "{email}", "password", "{password}")

	// CRUD for notes
	mx.HandleFunc("/create/note", handlers.CreateNoteHandler).Methods("POST").Queries("note", "{note}", "token", "{token}")
	mx.HandleFunc("/read/notes", handlers.ReadNotesHandler).Methods("GET").Queries("page", "{page}", "token", "{token}")
	mx.HandleFunc("/update/note", handlers.UpdateNotesHandler).Methods("PUT").Queries("id", "{id}", "note", "{note}", "token", "{token}")
	mx.HandleFunc("/delete/note", handlers.DeleteNoteHandler).Methods("DELETE").Queries("id", "{id}", "token", "{token}")

	return mx
}
