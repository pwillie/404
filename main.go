package main

import (
	"net/http"
	"log"
)

func healthcheck_handler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}

func default_handler(res http.ResponseWriter, req *http.Request) {
	log.Printf("not found: %s", req.URL.RequestURI())
	res.WriteHeader(http.StatusNotFound)
}

func main() {
	http.HandleFunc("/status", healthcheck_handler)
	http.HandleFunc("/", default_handler)
	http.ListenAndServe(":8080", nil)
}
