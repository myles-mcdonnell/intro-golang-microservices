package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var remoteServerPort string
var remoteServerName string

func main() {
	remoteServerPort = os.Getenv("REMOTE_SERVER_PORT")
	remoteServerName = os.Getenv("REMOTE_SERVER_NAME")
	fmt.Printf("REMOTE_SERVER_PORT %v\r\n", remoteServerPort)
	fmt.Printf("REMOTE_SERVER_NAME %v\r\n", remoteServerName)
	http.HandleFunc("/proxy", Proxy)
	http.HandleFunc("/hello", HelloServer)
	fmt.Println("Starting echoserver on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\r\n", r.Method,  r.URL.Path[1:])
	fmt.Fprint(w, "Hello")
}

func Proxy(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("PROXY %v\r\n", r.Method)
	proxyUrl := fmt.Sprintf("http://%v:%v/hello", remoteServerName, remoteServerPort)
	fmt.Printf("URL %v\r\n", proxyUrl)
	client := http.Client{
		Timeout:5*time.Second}

	req, err := http.NewRequest(
		r.Method,
		proxyUrl,
		bytes.NewBuffer([]byte{}))
	if err!=nil {
		fmt.Print(err.Error())
		w.WriteHeader(500)
		return
	}

	res, err := client.Do(req)
	if err!=nil {
		fmt.Print(err.Error())
		w.WriteHeader(500)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err!=nil {
		fmt.Print(err.Error())
		w.WriteHeader(500)
		return
	}

	w.Write(body)
}