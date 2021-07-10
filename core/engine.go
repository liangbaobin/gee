package gee

import (
	"log"
	"net/http"
	"strings"
)

// Engine is the uni handler for all requests
type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup // store all groups
}

func newEngine() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)

	c.handlers = middlewares
	engine.router.Handle(c)

}

func (engine *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, engine))
}

// Default use Logger() & Recovery middlewares
func Default() *Engine {
	engine := newEngine()
	engine.Use(Logger(), Recovery())
	return engine
}
