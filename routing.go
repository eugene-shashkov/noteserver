package noteserver

import (
	"noteserver/handlers"

	"github.com/gorilla/mux"
)

// Routing func is a startpoint of routing mechanism
func (srv Server) Routing() *mux.Router {

	mx := mux.NewRouter()
	mx.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	mx.HandleFunc("/api/registration", handlers.RegistrationHandler).Methods("POST").Queries("email", "{email}", "password", "{password}", "name", "{name}")
	mx.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST").Queries("email", "{email}", "password", "{password}")

	// CRUD for notes
	mx.HandleFunc("/api/create/note", handlers.CreateNoteHandler).Methods("POST").Queries("note", "{note}", "token", "{token}")
	mx.HandleFunc("/api/read/notes", handlers.ReadNotesHandler).Methods("GET").Queries("page", "{page}", "token", "{token}")
	mx.HandleFunc("/api/update/note", handlers.UpdateNotesHandler).Methods("PUT").Queries("id", "{id}", "note", "{note}", "token", "{token}")
	mx.HandleFunc("/api/delete/note", handlers.DeleteNoteHandler).Methods("DELETE").Queries("id", "{id}", "token", "{token}")

	return mx
}
