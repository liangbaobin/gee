package engine

import (
	"gee/context"
	"gee/router"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct {
	Router *router.Router
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := context.NewContext(w, req)
	engine.Router.Handle(c)
}

func New() (engine *Engine) {
	return &Engine{Router: router.New()}
}

func (engine *Engine) AddRouter(method string, pattern string, handler router.HandlerFunc) {
	engine.Router.AddRouter(method, pattern, handler)
}

func (engine *Engine) Run() {
	log.Fatal(http.ListenAndServe(":9999", engine))
}
