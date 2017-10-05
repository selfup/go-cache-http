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

// WriteToState is a safe way of updating state by forcing a Mutex Lock
func WriteToState(id string, lid string, state map[string][]string) {
	mutex.Lock()
	oldIds := state[lid]
	WriteToLidIds(oldIds, id, lid, state)
	mutex.Unlock()
}

// WriteToLidIds makes sure to not write to state under certain conditions
func WriteToLidIds(oldIds []string, id string, lid string, state map[string][]string) {
	var hasID bool

	for _, v := range oldIds {
		if v == id {
			hasID = true
			break // exit for loop as soon as we find a match
		} else {
			continue
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

	WriteToState(id, lid, state)

	fmt.Fprintf(w, "%s", state[lid])
}

func main() {
	http.HandleFunc("/", fetchCacheOrUpdate)
	http.ListenAndServe(PORT, nil)
}
