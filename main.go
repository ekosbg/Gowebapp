/*
Part of ChitChat
Ref: Go Web Programming Book
Part of Research
*/
package main

import (
	"net/http"
)

func main() {
	mux :+ http.NewServerMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}