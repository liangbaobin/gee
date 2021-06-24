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

func NewEngine() (engine *Engine) {
	engine2 := &Engine{router: newRouter()}
	engine2.RouterGroup = &RouterGroup{engine: engine2}
	engine2.groups = []*RouterGroup{engine2.RouterGroup}
	return engine2
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.Handle(c)
}

func (engine *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, engine))
}
