package main

import (
	"net/http"

	"trafilea/hdls"
	"trafilea/repos"
	"trafilea/srvs"

	"github.com/gorilla/mux"
)

func main() {
	repo := repos.New()
	srv := srvs.New(repo)
	hdl := hdls.New(srv)

	r := mux.NewRouter()
	r.HandleFunc("/numbers/{id:[0-9]+}", hdl.GetNumber).Methods("GET")
	r.HandleFunc("/numbers", hdl.GetNumbers).Methods("GET")
	r.HandleFunc("/numbers", hdl.PostNumber).Methods("POST")

	http.ListenAndServe(":8080", r)
}
