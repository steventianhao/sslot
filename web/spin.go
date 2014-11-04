package web

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/mux"
	"github.com/landjur/go-decimal"
	"log"
	"net/http"
	"sslot/games"
	"strconv"
	"time"
)

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
	if !games.ShowGame(gamename) {
		http.NotFound(w, r)
		return
	}

	// check if user authenticated
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		log.Println("connect to redis error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	//if user authed, then get the username, otherwise use session id as username
	username, _, err := GetUserName(conn, r)
	if err != nil {
		log.Println("get user name from redis error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if find user played this game before, then restore the state
	// otherwise just return a empty spin back to client

	history, err := games.RestoreSpinHistory(conn, username, gamename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		writeJson(w, r, history)
	}

}

func GetUserName(conn *redis.Client, r *http.Request) (string, bool, error) {
	sid, hash := AuthHash(r)
	reply := conn.Cmd("HGET", hash, sid)
	if isNil(reply) {
		return sid, true, nil
	}
	username, err := reply.Str()
	if err != nil {
		return "", true, err
	}
	if username == "" {
		return sid, true, nil
	}
	return username, false, nil
}

func FreeSpinGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gamename := vars["game"]
	if !games.ShowGame(gamename) {
		http.NotFound(w, r)
		return
	}
	// check if user authenticated
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		log.Println("connect to redis error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// //if user authed, then get the username, otherwise use session id as username
	username, _, err := GetUserName(conn, r)
	if err != nil {
		log.Println("get user name from redis error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result, err := games.PlayerFreeSpin(gamename, username); err != nil {
		log.Println("player free spin error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		writeJson(w, r, result)
	}
}

func NormalSpinGame(w http.ResponseWriter, r *http.Request) {
	//if game name given not right, return directly

	vars := mux.Vars(r)
	gamename := vars["game"]
	if !games.ShowGame(gamename) {
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
		log.Println("connect to redis error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// //if user authed, then get the username, otherwise use session id as username
	username, _, err := GetUserName(conn, r)
	if err != nil {
		log.Println("get user name from redis error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result, err := games.PlayerMainSpin(gamename, username, newlines, newbet); err != nil {
		log.Println("player main spin error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		writeJson(w, r, result)
	}
}
