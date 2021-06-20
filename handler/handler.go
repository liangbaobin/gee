package handler

import (
	"fmt"
	"gee/context"
	"log"
)

// IndexHandler handler echoes r.URL.Path
func IndexHandler(c *context.Context) {
	_, err := fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Path)
	if err != nil {
		log.Panic(err)
		return
	}
}

// HelloHandler handler echoes r.URL.Header
func HelloHandler(c *context.Context) {
	for k, v := range c.Req.Header {
		_, err := fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
		if err != nil {
			log.Panic(err)
			return
		}
	}
}
