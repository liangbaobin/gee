package gee

import (
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup // store all groups
}

func NewEngine() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.Handle(c)
}

func (engine *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, engine))
}
