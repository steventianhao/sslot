package web

import (
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var users = InitUsers()

type User struct {
	Id        int64
	Name      string
	anonymous bool
}

func NewUser(id int64, name string) *User {
	return &User{id, name, false}
}

func NewAnonymousUser(randomName string) *User {
	return &User{0, randomName, true}
}

func InitUsers() map[string]*User {
	us := make(map[string]*User)
	us["simon"] = NewUser(1, "simon")
	us["valor"] = NewUser(2, "valor")
	return us
}

func AuthUser(name string, password string) (*User, bool) {
	user, ok := users[name]
	return user, ok
}

func Auth(w http.ResponseWriter, r *http.Request) {
	sid, hash := AuthHash(r)
	vars := mux.Vars(r)
	name := vars["name"]
	password := vars["password"]
	u, found := AuthUser(name, password)
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
