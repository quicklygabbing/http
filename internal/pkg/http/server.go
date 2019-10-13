package http

import (
	"github.com/quicklygabbing/http/pkg/http/route"
)

// Server is implements http server workflow.
type Server interface {
	// Creates http server.
	Start(routes []route.Routes) error
	// Removes http server.
	Stop() error
}
