package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sslot/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/auth/{name}/{password}", web.Auth).Methods("GET")
	s := r.PathPrefix("/game").Subrouter()
	s.HandleFunc("/{game}/show", web.ShowGame).Methods("GET")
	s.HandleFunc("/{game}/spin", web.FreeSpinGame).Methods("GET")
	s.HandleFunc("/{game}/spin/{lines:[1-9][0-9]*}/{bet}", web.NormalSpinGame).Methods("GET")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("client/src"))))
	http.Handle("/", &web.SetSessionIfMissing{r})
	log.Println("Listening on 5555")
	http.ListenAndServe(":5555", nil)
}
