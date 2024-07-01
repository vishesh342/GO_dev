package main

import (
	"net/http"
	"log"
	"microservice/handler"
)
func main(){
	logger := log.New(log.Writer(), "logger: ", log.LstdFlags)
	sh := handler.NewServerLogger(logger)
	sm := http.NewServeMux()
	sm.Handle("/", sh)

	http.ListenAndServe(":9090",sm)
}