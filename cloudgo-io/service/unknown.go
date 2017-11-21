package service

import (
	"net/http"
)

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "502 Not Implemented", http.StatusNotImplemented)
}

func NotImplementedHandler() http.Handler {
	return http.HandlerFunc(NotImplemented)
}
