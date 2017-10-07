package main

import (
	"encoding/json"
	"net/http"
	"os"

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
	decoder := json.NewDecoder(r.Body)

	var incoming struct {
		Key     string
		Data    string
		Unix    int64
		Expires int64
	}

	err := decoder.Decode(&incoming)
	if err != nil {
		http.Error(w, "failed to parse JSON", 500)
		return
	}

	lib.WriteToState(
		incoming.Key,
		incoming.Data,
		incoming.Unix,
		incoming.Expires,
		state,
	)

	outgoing, err := json.Marshal(incoming)
	if err != nil {
		http.Error(w, "failed to stringify JSON", 500)
		return
	}

	w.Write(outgoing)
}

func main() {
	port := definePort()

	http.HandleFunc("/", fetchCacheOrUpdate)
	http.ListenAndServe(port, nil)
}
