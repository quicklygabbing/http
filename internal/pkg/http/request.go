package http

import (
	"net/http"

	"github.com/quicklygabbing/http/internal/pkg/http/handle"
)

// Requests is implements router of http server.
type Request interface {
	// Returns response with error 404 code.
	NotFound() http.Handler
	// Returns response for not allowed methods (GET/POST etc).
	MethodNotAllowed() http.Handler
	// Returns shutdown response for panic error.
	PanicHandler(http.ResponseWriter, *http.Request, interface{})
}

type requests struct {
}

func NewRequests() Request {
	return &requests{}
}

func (r *requests) NotFound() http.Handler {
	return handle.NotFound()
}

func (r *requests) MethodNotAllowed() http.Handler {
	return handle.MethodNotAllowed()
}

func (r *requests) PanicHandler(http.ResponseWriter, *http.Request, interface{}) {
}
