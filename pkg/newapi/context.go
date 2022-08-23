package api

import (
	"net/http"

	xcontext "github.com/gorilla/context"
)

func GetStringOk(r *http.Request, key interface{}) (string, bool) {
	//val, ok := xcontext.GetOk(r, key)
	val := r.Context().Value(key)
	s, ok := val.(string)
	return s, ok
}

func GetXContextStringOk(r *http.Request, key interface{}) (string, bool) {
	val, ok := xcontext.GetOk(r, key)
	if !ok {
		return "", false
	}
	s, ok := val.(string)
	return s, ok
}
