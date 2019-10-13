package handle

import (
	"net/http"
)

type notFound struct {
}

func NotFound() http.Handler {
	return &notFound{}
}

func (n *notFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
