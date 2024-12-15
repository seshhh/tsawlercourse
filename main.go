package main

import (
	"net/http"

	"github.com/seshhh/tsawlercourse/handlers"
)


func main(){
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/goiz", handlers.Goiz)
	http.ListenAndServe(":8080", nil)
}

