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
	/*
		In main.go, you first create a multiplexer, 
		the piece of code that redirects a request to a handler. The net/http standard 
		library provides a default multiplexer that can be created by 
		calling the NewServeMux function:
	*/
	mux :+ http.NewServerMux()
	/*
		Besides redirecting to the appropriate handler, you can use the multiplexer to serve
		static files. To do this, you use the FileServer function to create a handler that will
		serve files from a given directory. Then you pass the handler to the Handle function of
		the multiplexer. You use the StripPrefix function to remove the given prefix from
		the request URL ’s path.
	*/
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	//To redirect the root URL to a handler function, you use the HandleFunc function:
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}