package game

import (
	"encoding/base64"
	"errors"
	"github.com/gorilla/securecookie"
)

const (
	SESSION_ID = "SID"
)

func RandomString() (string, error) {
	if key := securecookie.GenerateRandomKey(128); key == nil {
		return "", errors.New("generate random key failed")
	} else {
		return base64.StdEncoding.EncodeToString(key), nil
	}
}
