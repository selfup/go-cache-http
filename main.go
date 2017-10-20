package main

import (
	"net/http"

	"github.com/selfup/gocrashttp/handlers"
	"github.com/selfup/gocrashttp/lib"
)

var (
	state = make(map[string]*lib.CacheData)
)

func main() {
	http.HandleFunc("/", handlers.FetchCacheOrUpdate(state))
	http.ListenAndServe(handlers.DefinePort(), nil)
}
