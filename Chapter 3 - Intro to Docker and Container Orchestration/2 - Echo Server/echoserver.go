package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("Starting echoserver on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\r\n", r.Method,  r.URL.Path[1:])
	fmt.Fprintln(w, fmt.Sprintf("Hello, %s!", r.URL.Path[1:]))
}