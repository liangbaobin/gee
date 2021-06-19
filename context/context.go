package context

import (
	"net/http"
)

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string

	// response info
	StatusCode int
}

func NewContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: writer,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}

func (c Context) PostForm(key string) string{
	return c.Req.FormValue(key)
}


