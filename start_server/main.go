package main

import (
	"log"
	"net/http"
	"noteserver"
)

func main() {
	sv := noteserver.Server{}
	sv.CreateServer(&sv)
	routing := sv.Routing()
	log.Fatal(http.ListenAndServe(":5430", routing))
}
