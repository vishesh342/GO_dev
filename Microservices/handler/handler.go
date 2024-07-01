package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ServerHandler is an HTTP handler that responds with a "Hello" message.
type ServerHandler struct {
	logger *log.Logger
}

// NewServerLogger returns a new ServerHandler instance with the given logger.
func NewServerLogger(logger *log.Logger) *ServerHandler {
	return &ServerHandler{logger}
}

// ServeHTTP implements the http.Handler interface.
func (sh *ServerHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	sh.logger.Println("Handle Hello requests")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err!= nil {
		sh.logger.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Hello %s", b)
}
