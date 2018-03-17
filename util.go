package main

import (
	"errors"
	"net/http"
)

// func method(引数)(戻り値、複数の場合)
func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	// リクエストからクッキーを取り出す
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
