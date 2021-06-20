package router

import (
	"gee/context"
	"log"
	"net/http"
)

type HandlerFunc func(c *context.Context)

type Router struct {
	Handlers map[string]HandlerFunc
}

func New() (router *Router) {
	return &Router{Handlers: make(map[string]HandlerFunc)}
}

func (r *Router) AddRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.Handlers[key] = handler
}

func (r *Router) Handle(c *context.Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.Handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
