package main

import (
	"net/http"
	"log"
)
func main(){
	log := log.New(log.Writer(), "logger: ", log.LstdFlags)
	// sh := handler.NewHello(*log)

	http.ListenAndServe(":9090",nil)
}