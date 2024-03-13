package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	path    string
	method  string
	handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      []Handler
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]Handler, 0),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, Handler{
		path:    path,
		method:  method,
		handler: handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		s.Router.Method(handler.method, handler.path, handler.handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
