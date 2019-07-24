package main

import (
	"log"
	"net/http"
)

func main() {
	routing := Routing()
	log.Fatal(http.ListenAndServe(":3429", routing))
}
