package main

import (
	"log"
	"net/http"
)

// type HandlerFunc func(ResponseWriter, *Request)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
