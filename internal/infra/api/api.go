package api

import (
	"fmt"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/ianbrito/fr-cotacao/docs"
)

type WebServer struct {
	Router   *chi.Mux
	HttpPort string
	Handlers map[string]http.HandlerFunc
}

func NewWebServer(httpPort string) *WebServer {

	if httpPort == "" {
		httpPort = "80"
	}

	return &WebServer{
		Router:   chi.NewRouter(),
		HttpPort: httpPort,
		Handlers: make(map[string]http.HandlerFunc),
	}
}

func (w *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	w.Handlers[path] = handler
}

func (w *WebServer) Run() {
	var addr = fmt.Sprintf(":%s", w.HttpPort)

	w.Router.Use(middleware.Logger)
	w.Router.Use(render.SetContentType(render.ContentTypeJSON))
	w.Router.Mount("/swagger", httpSwagger.WrapHandler)
	for path, handler := range w.Handlers {
		w.Router.Handle(path, handler)
	}

	fmt.Printf("Listening on 0.0.0.0:%s\n", w.HttpPort)
	if err := http.ListenAndServe(addr, w.Router); err != nil {
		panic(err.Error())
	}
}
