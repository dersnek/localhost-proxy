package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var defaultHost = ""

func main() {
	port := ""

	fmt.Println("Which port to send requests to?")
	fmt.Scanln(&port)

	defaultHost = "http://127.0.0.1:" + port

	http.HandleFunc("/", handleReq)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleReq(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Request received: " + req.Host + req.URL.Path)
	fmt.Println("Sending to: " + defaultHost + req.URL.Path)

	serveReverseProxy(res, req)
}

func serveReverseProxy(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(defaultHost)

	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ServeHTTP(res, req)
	fmt.Println("Completed")
}
