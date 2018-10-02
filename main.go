package main

import (
	"log"      
	"net/http"
	"simple_HTTP_server/handlers"
)



func main() {
	http.HandleFunc("/date", handlers.Date) 
	http.HandleFunc("/wait", handlers.Wait)
	err := http.ListenAndServe(":9000", nil) 
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
