package main

import (
	"gee/engine"
	"gee/handler"
)

func main() {
	gee := engine.New()
	gee.AddRouter("GET", "/", handler.IndexHandler)
	gee.AddRouter("GET", "/hello", handler.HelloHandler)
	gee.Run()
}
