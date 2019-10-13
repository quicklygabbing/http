package handle

import (
	"net/http"
)

type methodNotAllowed struct {
}

func MethodNotAllowed() http.Handler {
	return &methodNotAllowed{}
}

func (m *methodNotAllowed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
