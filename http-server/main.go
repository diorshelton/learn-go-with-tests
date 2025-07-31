package main

import (
	"log"
	"net/http"
)

type InMemoryPlayStore struct{}

func (i *InMemoryPlayStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
