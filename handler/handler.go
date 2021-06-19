package handler

import (
	"fmt"
	"log"
	"net/http"
)

// IndexHandler handler echoes r.URL.Path
func IndexHandler(writer http.ResponseWriter, request *http.Request){
	_, err := fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	if err != nil {
		log.Panic(err)
		return
	}
}

// HelloHandler handler echoes r.URL.Header
func HelloHandler(writer http.ResponseWriter, request *http.Request){
	for k, v := range request.Header {
		_, err := fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		if err != nil {
			log.Panic(err)
			return
		}
	}
}
