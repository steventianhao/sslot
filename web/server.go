package main

import (
	"encoding/json"
	"fmt"
	//"github.com/gorilla/handlers"
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/mux"
	"github.com/landjur/go-decimal"
	"log"
	"net/http"
	"sslot/web/game"
	"strconv"
	"time"
)

type SetSessionIfMissing struct {
	http.Handler
}

func (f *SetSessionIfMissing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie(game.SESSION_ID); err != nil {
		if sid, err := game.RandomString(); err == nil {
			cookie := &http.Cookie{Name: game.SESSION_ID, Value: sid, Path: "/"}
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

func writeJson(w http.ResponseWriter, r *http.Request, obj interface{}) {
	if bytes, err := json.Marshal(obj); err != nil {
		http.NotFound(w, r)
	} else {
		w.Write(bytes)
	}
}

func AuthHash(r *http.Request) (string, string) {
	sess, _ := r.Cookie(game.SESSION_ID)
	sid := sess.Value
	return sid, game.RedisHashKey("auths", sid)
}

func UserHash(username string) string {
	return fmt.Sprint("user:", username)
}

func GameFieldLines(gamename string) string {
	return fmt.Sprint("game_", gamename, "_lines")
}
func GameFieldBet(gamename string) string {
	return fmt.Sprint("game_", gamename, "_bet")
}

func GameFieldFeatures(gamename string) string {
	return fmt.Sprint("game_", gamename, "_features")
}

func ShowGame(w http.ResponseWriter, r *http.Request) {
	// if game name is not valid, return directly
	vars := mux.Vars(r)
	gamename := vars["game"]
	if !game.ShowGame(gamename) {
		http.NotFound(w, r)
		return
	}

	// check if user authenticated
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	//if user authed, then get the username, otherwise use session id as username
	username, _, err := GetUserName(conn, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if find user played this game before, then restore the state
	// otherwise just return a empty spin back to client

	spin, err := RestoreSpin(conn, username, gamename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		writeJson(w, r, spin)
	}

}

func GetUserName(conn *redis.Client, r *http.Request) (string, bool, error) {
	sid, hash := AuthHash(r)
	username, err := conn.Cmd("HGET", hash, sid).Str()
	if err != nil {
		return "", false, err
	}
	if username == "" {
		return sid, true, nil
	}
	return username, false, nil
}

func RestoreSpin(conn *redis.Client, username, gamename string) (*game.Spin, error) {
	key := UserHash(username)
	fmt.Println("key for restore spin:", key)
	res, err := conn.Cmd("HMGET", key, GameFieldLines(gamename), GameFieldBet(gamename), GameFieldFeatures(gamename)).List()
	if err != nil {
		return nil, err
	}
	strLine, strBet, strFeatures := res[0], res[1], res[2]
	fmt.Println("restore values", strLine, strBet, strFeatures)
	if strLine != "" && strBet != "" && strFeatures != "" {
		lines, err := strconv.Atoi(strLine)
		if err != nil {
			return nil, err
		}
		featrues, err := strconv.Atoi(strFeatures)
		if err != nil {
			return nil, err
		}
		return game.CacheSpin(gamename, lines, strBet, featrues), nil
	} else {
		return game.FreshSpin(gamename), nil
	}
}

func FreeSpinGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gamename := vars["game"]
	if !game.ShowGame(gamename) {
		http.NotFound(w, r)
		return
	}
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	username, _, err := GetUserName(conn, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	spin, err := RestoreSpin(conn, username, gamename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if spin.Features < 1 {
		http.NotFound(w, r)
		return
	}

	bet, err := decimal.Parse(spin.Bet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	/*
		decrease the features, if >=0 return 1, if <0, set to 0 return 0
		eval "if redis.call('HINCRBY',KEYS[1],KEYS[2],-1)>=0 then return 1 else redis.call('HSET',KEYS[1],KEYS[2],0) return 0 end" 2 myhash field1
	*/
	luaScript := "local v = redis.call('HINCRBY',KEYS[1],KEYS[2],-1) if v>=0 then return v else redis.call('HSET',KEYS[1],KEYS[2],0) return -1 end"
	n, err := conn.Cmd("eval", luaScript, 2, UserHash(username), GameFieldFeatures(gamename)).Int()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if n >= 0 {
		if spinGame, err := game.FreeSpinGame(gamename, spin.Lines, bet); err != nil {
			http.NotFound(w, r)
		} else {
			n2, err := conn.Cmd("HINCRBY", UserHash(username), GameFieldFeatures(gamename), spinGame.Features).Int()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				spinGame.Features = n2
				writeJson(w, r, spinGame)
			}
		}
	} else {
		http.NotFound(w, r)
	}
}

func NormalSpinGame(w http.ResponseWriter, r *http.Request) {
	//if game name given not right, return directly

	vars := mux.Vars(r)
	gamename := vars["game"]
	if !game.ShowGame(gamename) {
		http.NotFound(w, r)
		return
	}
	newlines, err := strconv.Atoi(vars["lines"])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	newbet, err := decimal.Parse(vars["bet"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// check if user authenticated
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	//if user authed, then get the username, otherwise use session id as username

	username, _, err := GetUserName(conn, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//restore the previous spin
	spin, err := RestoreSpin(conn, username, gamename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if spin.Features > 0 {
		http.NotFound(w, r)
		return
	}

	if spinGame, err := game.SpinGame(gamename, newlines, newbet); err != nil {
		http.NotFound(w, r)
	} else {
		conn.Cmd("HMSET", UserHash(username), GameFieldLines(gamename), newlines, GameFieldBet(gamename), newbet.String(), GameFieldFeatures(gamename), spinGame.Features).List()
		writeJson(w, r, spinGame)
	}
}

func main() {
	game.InitUsers()

	r := mux.NewRouter()
	r.HandleFunc("/auth/{name}/{password}", Auth).Methods("GET")
	s := r.PathPrefix("/game").Subrouter()
	s.HandleFunc("/{game}/show", ShowGame).Methods("GET")
	s.HandleFunc("/{game}/spin", FreeSpinGame).Methods("GET")
	s.HandleFunc("/{game}/spin/{lines:[1-9][0-9]*}/{bet}", NormalSpinGame).Methods("GET")
	r.HandleFunc("/", Index).Methods("GET")
	http.Handle("/", &SetSessionIfMissing{r})
	log.Println("Listening on 5555")
	http.ListenAndServe(":5555", nil)
}
