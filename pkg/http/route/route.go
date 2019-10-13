package route

import (
	"github.com/julienschmidt/httprouter"
)

const (
	GET  = `GET`
	POST = `POST`
	PUT  = `PUT`
)

type Routes struct {
	Route   string
	Methods []string
	Handle  httprouter.Handle
}
