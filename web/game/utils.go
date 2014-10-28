package game

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gorilla/securecookie"
	"hash/fnv"
)

const (
	REDIS_HASH_MAX = 1000
	SESSION_ID     = "SID"
)

func Hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
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
