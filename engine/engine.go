package engine

import (
	"gee/router"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct{
	Router *router.Router
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := engine.Router.Handlers[req.URL.Path]; ok{
		handler(w, req)
	}
}

func New()(engine *Engine){
	return &Engine{Router: router.New()}
}

func (engine *Engine) AddRouter(path string, handler router.HandlerFunc){
	engine.Router.Handlers[path] = handler
}

func (engine *Engine) Run(){
	log.Fatal(http.ListenAndServe(":9999", engine))
}