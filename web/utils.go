package web

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/context"
	"github.com/gorilla/securecookie"
	"hash/fnv"
	"net/http"
)

const (
	REDIS_HASH_MAX  = 1000
	SESSION_ID      = "SID"
	CTX_SESSION_KEY = 0
)

type SetSessionIfMissing struct {
	http.Handler
}

func (f *SetSessionIfMissing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c, err := r.Cookie(SESSION_ID); err != nil {
		if sid, err := RandomString(); err == nil {
			cookie := &http.Cookie{Name: SESSION_ID, Value: sid, Path: "/"}
			http.SetCookie(w, cookie)
			context.Set(r, CTX_SESSION_KEY, sid)
			f.Handler.ServeHTTP(w, r)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		context.Set(r, CTX_SESSION_KEY, c.Value)
		f.Handler.ServeHTTP(w, r)
	}
}

func AuthHash(r *http.Request) (string, string) {
	sid := context.Get(r, CTX_SESSION_KEY).(string)
	return sid, RedisHashKey("auths", sid)
}

func writeJson(w http.ResponseWriter, r *http.Request, obj interface{}) {
	if bytes, err := json.Marshal(obj); err != nil {
		http.NotFound(w, r)
	} else {
		w.Write(bytes)
	}
}

func Hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func isNil(r *redis.Reply) bool {
	return r.Type == redis.NilReply
}

func RedisHashKey(prefix, value string) string {
	return fmt.Sprint(prefix, ":", Hash(value)/REDIS_HASH_MAX)
}

func RandomString() (string, error) {
	if key := securecookie.GenerateRandomKey(128); key == nil {
		return "", errors.New("generate random key failed")
	} else {
		return base64.StdEncoding.EncodeToString(key), nil
	}
}
