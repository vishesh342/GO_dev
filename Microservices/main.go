package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
		log.Println("Server Listening..")
		body, err := ioutil.ReadAll(r.Body)
		if err!=nil{
			http.Error(rw,"Error Occured...",http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw,"Data->%s",body)

	})
	http.HandleFunc("/goodbuy",func(http.ResponseWriter,*http.Request){
		log.Println("Good Bye....")
	})

	http.ListenAndServe(":9090",nil)
}