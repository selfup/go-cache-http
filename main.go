package main

import (
	"fmt"
	"net/http"

	lib "github.com/selfup/go_cache_http/lib"
)

const (
	// PORT is the server port
	PORT = ":8080"
)

var (
	state = make(map[string][]string)
)

func fetchCacheOrUpdate(w http.ResponseWriter, r *http.Request) {
	lid := r.FormValue("lid")
	id := r.FormValue("id")

	lib.WriteToState(id, lid, state)

	fmt.Fprintf(w, "%s", state[lid])
}

func main() {
	http.HandleFunc("/", fetchCacheOrUpdate)
	http.ListenAndServe(PORT, nil)
}
