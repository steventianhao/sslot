package web

import (
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/mux"
	"net/http"
	"sslot/web/game"
	"time"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	sid, hash := AuthHash(r)
	vars := mux.Vars(r)
	name := vars["name"]
	password := vars["password"]
	u, found := game.AuthUser(name, password)
	if !found {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	_, err = conn.Cmd("HSETNX", hash, sid, u.Name).Int()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
