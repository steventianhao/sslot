package web

import (
	"encoding/json"
	"net/http"
	"sslot/web/game"
)

func AuthHash(r *http.Request) (string, string) {
	sess, _ := r.Cookie(game.SESSION_ID)
	sid := sess.Value
	return sid, game.RedisHashKey("auths", sid)
}

func writeJson(w http.ResponseWriter, r *http.Request, obj interface{}) {
	if bytes, err := json.Marshal(obj); err != nil {
		http.NotFound(w, r)
	} else {
		w.Write(bytes)
	}
}
