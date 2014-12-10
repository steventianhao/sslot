package web

import (
	"errors"
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/mux"
	"github.com/landjur/go-decimal"
	"log"
	"net/http"
	"sslot/engine"
	"sslot/games"
	"strconv"
	"time"
)

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

func GetUserNameFromRedis(r *http.Request) (string, error) {
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	username, _, err := GetUserName(conn, r)
	if err != nil {
		return "", err
	}
	return username, nil
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

func NormalSpinGame(w http.ResponseWriter, r *http.Request) {
	sp, err := validateNormalSpinReq(r)
	SpinGame(w, r, sp, err)
}

func FreeSpinGame(w http.ResponseWriter, r *http.Request) {
	sp, err := validateFreeSpinReq(r)
	SpinGame(w, r, sp, err)
}

type SpinAction interface {
	gameName() string
	spin(username string) (*engine.SpinResult, error)
}

type FreeSpinParams struct {
	game string
}

func (sp FreeSpinParams) gameName() string {
	return sp.game
}

func (sp FreeSpinParams) spin(username string) (*engine.SpinResult, error) {
	return games.PlayerFreeSpin(sp.game, username)
}

type NormalSpinParams struct {
	game  string
	lines int
	bet   *decimal.Decimal
}

func (sp NormalSpinParams) gameName() string {
	return sp.game
}

func (sp NormalSpinParams) spin(username string) (*engine.SpinResult, error) {
	return games.PlayerMainSpin(sp.game, username, sp.lines, sp.bet)
}

func validateFreeSpinReq(r *http.Request) (*FreeSpinParams, error) {
	vars := mux.Vars(r)
	game, ok := vars["game"]
	if !ok {
		return nil, errors.New("Param game required")
	}
	return &FreeSpinParams{game}, nil
}

func validateNormalSpinReq(r *http.Request) (*NormalSpinParams, error) {
	vars := mux.Vars(r)

	game, ok := vars["game"]
	if !ok {
		return nil, errors.New("Param game required")
	}

	sLines, ok := vars["lines"]
	if !ok {
		return nil, errors.New("Param lines required")
	}

	lines, err := strconv.Atoi(sLines)
	if err != nil {
		return nil, errors.New("Param lines is not number")
	}

	sBet, ok := vars["bet"]
	if !ok {
		return nil, errors.New("Param bet required")
	}

	bet, err := decimal.Parse(sBet)
	if err != nil {
		return nil, errors.New("Param bet is not decimal")
	}

	return &NormalSpinParams{game, lines, bet}, nil
}

func SpinGame(w http.ResponseWriter, r *http.Request, sp SpinAction, valiateErr error) {
	if valiateErr != nil {
		http.NotFound(w, r)
		return
	}

	if !games.ShowGame(sp.gameName()) {
		http.NotFound(w, r)
		return
	}

	username, err := GetUserNameFromRedis(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result, err := sp.spin(username); err != nil {
		log.Println("player main spin error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		writeJson(w, r, result)
	}
}
