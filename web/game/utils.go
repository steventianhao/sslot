package game

import (
	"fmt"
	"hash/fnv"
)

const (
	REDIS_HASH_MAX = 1000
)

func Hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func RedisHashKey(prefix, value string) string {
	return fmt.Sprint(prefix, ":", Hash(value)/REDIS_HASH_MAX)
}
