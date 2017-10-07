package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	lib "github.com/selfup/go-cache-http/lib"
)

var (
	state = make(map[string]*lib.CacheData)
)

func definePort() string {
	portEnv := os.Getenv("PORT")

	if portEnv != "" {
		return ":" + portEnv
	}

	return ":8080"
}

func fetchCacheOrUpdate(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	data := r.FormValue("data")
	unixString := r.FormValue("unix")

	unix, err := strconv.ParseInt(unixString, 0, 64)
	if err != nil {
		panic(err)
	}

	lib.WriteToState(key, data, unix, state)

	fmt.Fprintf(w, "%v", state[key])
}

func main() {
	port := definePort()

	http.HandleFunc("/", fetchCacheOrUpdate)
	http.ListenAndServe(port, nil)
}
