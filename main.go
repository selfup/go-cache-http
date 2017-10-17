package main

import (
	"encoding/json"
	"net/http"
	"os"

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

// FetchCacheOrUpdate is the main handler
func FetchCacheOrUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "this endpoint only supports POST requests", 405)
		return
	}

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

	state[incoming.Key].ExpirationValidator()
	if state[incoming.Key].Valid == false {
		delete(state, incoming.Key)
	}

	outgoing, err := json.Marshal(state[incoming.Key])
	if err != nil {
		http.Error(w, "failed to stringify JSON", 500)
		return
	}

	w.Write(outgoing)
}

func main() {
	http.HandleFunc("/", FetchCacheOrUpdate)
	http.ListenAndServe(DefinePort(), nil)
}
