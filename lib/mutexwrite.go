package lib

import (
	"fmt"
	"sync"
)

var (
	mutex = &sync.Mutex{}
)

// WriteToState is a safe way of updating state by forcing a Mutex Lock
func WriteToState(
	lid string,
	data string,
	unix int64,
	state map[string]*CacheData,
) {
	mutex.Lock()
	WriteToLidIds(lid, data, unix, state)
	mutex.Unlock()
}

// WriteToLidIds makes sure to not write to state under certain conditions
func WriteToLidIds(
	lid string,
	data string,
	unix int64,
	state map[string]*CacheData,
) {
	if state[lid] != nil {
		fmt.Println("wow")
	} else {
		state[lid] = NewCacheData(data, unix)
	}
}
