package web

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/securecookie"
	"hash/fnv"
	"net/http"
)

const (
	REDIS_HASH_MAX = 1000
	SESSION_ID     = "SID"
)

func AuthHash(r *http.Request) (string, string) {
	sess, _ := r.Cookie(SESSION_ID)
	sid := sess.Value
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
