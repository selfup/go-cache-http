package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	// PORT is the server port
	PORT = ":8080"
)

var (
	state = make(map[string][]string)
	mutex = &sync.Mutex{}
)

// User is just an id and a unix timestamp object
type User struct {
	id   string
	unix int64
}

func writeToState(id string, lid string, state map[string][]string) {
	mutex.Lock()
	oldIds := state[lid]
	writeToLidIds(oldIds, id, lid, state)
	mutex.Unlock()
}

func writeToLidIds(oldIds []string, id string, lid string, state map[string][]string) {
	var hasID bool

	for _, v := range oldIds {
		if v == id {
			hasID = true
		}
	}

	if !hasID {
		ids := append(oldIds, id)
		state[lid] = ids
	}
}

func fetchCacheOrUpdate(w http.ResponseWriter, r *http.Request) {
	lid := r.FormValue("lid")
	id := r.FormValue("id")

	writeToState(id, lid, state)

	fmt.Fprintf(w, "%s", state[lid])
}

func main() {
	http.HandleFunc("/", fetchCacheOrUpdate)
	http.ListenAndServe(PORT, nil)
}
