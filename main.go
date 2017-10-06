package main

import (
	"fmt"
	"net/http"
	"os"

	lib "github.com/selfup/go-cache-http/lib"
)

var (
	port  = definePort()
	state = make(map[string][]string)
)

func definePort() string {
	portEnv := os.Getenv("PORT")

	if portEnv != "" {
		return portEnv
	}

	return ":8080"
}

func fetchCacheOrUpdate(w http.ResponseWriter, r *http.Request) {
	lid := r.FormValue("lid")
	id := r.FormValue("id")

	lib.WriteToState(id, lid, state)

	fmt.Fprintf(w, "%s", state[lid])
}

func main() {
	http.HandleFunc("/", fetchCacheOrUpdate)
	http.ListenAndServe(port, nil)
}
