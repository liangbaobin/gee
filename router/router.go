package router

import "net/http"

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Router struct {
	Handlers map[string]HandlerFunc
}

func New()(router *Router){
	return &Router{Handlers: make(map[string]HandlerFunc)}
}

