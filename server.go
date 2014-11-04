package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sslot/web"
)

type SetSessionIfMissing struct {
	http.Handler
}

func (f *SetSessionIfMissing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie(web.SESSION_ID); err != nil {
		if sid, err := web.RandomString(); err == nil {
			cookie := &http.Cookie{Name: web.SESSION_ID, Value: sid, Path: "/"}
			http.SetCookie(w, cookie)
			f.Handler.ServeHTTP(w, r)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		f.Handler.ServeHTTP(w, r)
	}
}

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

	http.Handle("/", &SetSessionIfMissing{r})
	log.Println("Listening on 5555")
	http.ListenAndServe(":5555", nil)
}
