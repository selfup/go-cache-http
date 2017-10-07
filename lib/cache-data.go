package lib

// CacheData is to be held as a value if state.
// It will provide the following features
// 1. It will store the data string
// 2. It will store a unix timestamp
// 3. It will check if the object itself has expired
// 4. IF the object has expired -> expose true for the valid prop
// 5. IF the data needs to be fresh, it will override internal props
type CacheData struct {
	// simple string or JSON/XML
	data string
	// use int64 as unix timestamps will have to change to int64 in 2038
	unix int64
	// bit flippin flag
	valid bool
}

// NewCacheData will take data and timestamp and default valid to true
func NewCacheData(data string, unix int64) *CacheData {
	return &CacheData{
		data:  data,
		unix:  unix,
		valid: true,
	}
}
