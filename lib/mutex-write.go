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
	WriteCacheValueData(lid, data, unix, state)
	mutex.Unlock()
}

// WriteCacheValueData makes sure to not write to state under certain conditions
func WriteCacheValueData(
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
