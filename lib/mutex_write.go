package lib

import (
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
	expires int64,
	state map[string]*CacheData,
) {
	mutex.Lock()
	WriteCacheValueData(lid, data, unix, expires, state)
	potentialKeyRemove(state, lid)
	mutex.Unlock()
}

// potentialKeyRemove checks to see if it needs to remove the k/v pair
// first by checking to see if the CacheData is valid or not
// if not valid - yank it!
func potentialKeyRemove(state map[string]*CacheData, lid string) {
	state[lid].ExpirationValidator()
	if state[lid].Valid == false {
		delete(state, lid)
	}
}

// WriteCacheValueData makes sure to not write to state under certain conditions
func WriteCacheValueData(
	lid string,
	data string,
	unix int64,
	expires int64,
	state map[string]*CacheData,
) {
	if state[lid] != nil {

	} else {
		state[lid] = NewCacheData(data, unix, expires)
	}
}
