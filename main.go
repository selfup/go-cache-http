package main

import (
	"net/http"
	"os"

	"github.com/selfup/gocrashttp/handlers"
	"github.com/selfup/gocrashttp/lib"
)

var (
	state = make(map[string]*lib.CacheData)
)

// DefinePort either grabs the PORT ENV or defines 8080 as the port
func DefinePort() string {
	portEnv := os.Getenv("PORT")

	if portEnv != "" {
		return ":" + portEnv
	}

	return ":8080"
}

func main() {
	http.HandleFunc("/", handlers.FetchCacheOrUpdate(state))
	http.ListenAndServe(DefinePort(), nil)
}
