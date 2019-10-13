package http

import (
	"net/http"

	internalHttp "github.com/quicklygabbing/http/internal/pkg/http"
	"github.com/quicklygabbing/http/pkg/http/route"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

type server struct {
	address *string
	request internalHttp.Request
	router  *httprouter.Router
	server  *http.Server
}

func NewServer(address *string) internalHttp.Server {
	return &server{address: address}
}

func (s *server) Start(routes []route.Routes) (err error) {
	err = s.setupServer(routes)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *server) Stop() (err error) {
	err = s.server.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	s.reset()

	return
}

func (s *server) setupServer(routes []route.Routes) (err error) {
	s.setupRequest()
	s.setupRouter(routes)
	s.server = &http.Server{Addr: *s.address, Handler: s.router}
	err = s.server.ListenAndServe()
	if err != nil {
		s.reset()
		return errors.WithStack(err)
	}
	return nil
}

func (s *server) setupRequest() {
	s.request = internalHttp.NewRequests()
}

func (s *server) setupRouter(routes []route.Routes) {
	s.router = &httprouter.Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
		NotFound:               s.request.NotFound(),
		MethodNotAllowed:       s.request.MethodNotAllowed(),
		PanicHandler:           s.request.PanicHandler,
	}

	s.route(routes)
}

func (s *server) route(routes []route.Routes) {
	for _, value := range routes {
		for _, f := range value.Methods {
			s.router.Handle(f, value.Route, value.Handle)
		}
	}
}

func (s *server) reset() {
	s.server = nil
	s.router = nil
	s.request = nil
}
