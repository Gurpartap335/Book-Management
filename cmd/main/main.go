package main

import (
	"log"
	"net/http"

	"github.com/Gurpartap335/Book-Management/pkg/routes"
	"github.com/gorilla/mux"
)

//create server
// where ever routes reside

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("locakhost:9010", r))
}
