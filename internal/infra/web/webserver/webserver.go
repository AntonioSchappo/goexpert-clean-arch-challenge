package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Handlers map[string]http.HandlerFunc
	Router   chi.Router
	Port     string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Handlers: make(map[string]http.HandlerFunc),
		Router:   chi.NewRouter(),
		Port:     port,
	}
}

func (ws *WebServer) Start() {
	ws.Router.Use(middleware.Logger)
	for path, handler := range ws.Handlers {
		ws.Router.Handle(path, handler)
	}
	http.ListenAndServe(ws.Port, ws.Router)
}

func (ws *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	ws.Handlers[path] = handler
}
