package main

import (
	"gee/engine"
	"gee/handler"
)

func main()  {
	gee := engine.New()
	gee.AddRouter("/", handler.IndexHandler)
	gee.AddRouter("/hello", handler.HelloHandler)
	gee.Run()
}
