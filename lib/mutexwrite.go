package lib

import (
	"sync"
)

var (
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
func WriteToLidIds(
	oldIds []string, id string, lid string, state map[string][]string,
) {
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
