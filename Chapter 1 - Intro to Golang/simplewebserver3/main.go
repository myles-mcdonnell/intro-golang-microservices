package main

import (
	"net/http"
	"github.com/myles-mcdonnell/simplewebserver2/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HelloServer)
	http.ListenAndServe(":8080", nil)
}
