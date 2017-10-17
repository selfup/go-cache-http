package lib

// CacheData is to be held as a value if state.
// It will provide the following features
// 1. It will store the data string
// 2. It will store a unix timestamp
// 3. It will store an experiation time limit
// 4. It will check if the object itself has expired
// 5. IF the object has expired -> expose true for the valid prop
// 6. IF the data needs to be fresh, it will override internal props
type CacheData struct {
	// simple string or JSON/XML
	data string
	// use int64 as unix timestamps will have to change to int64 in 2038
	unix int64
	// default to never otherwise invalidate when asked
	expires int64
	// bit flippin flag
	Valid bool
}

// NewCacheData will take data and timestamp and default valid to true
func NewCacheData(data string, unix int64, expires int64) *CacheData {
	return &CacheData{
		data:    data,
		unix:    unix,
		expires: expires,
		Valid:   true,
	}
}

// ExpirationValidator checks to see if a CacheData object is valid or not
func (c *CacheData) ExpirationValidator() {
	// if expires is not the default value
	// check to see if the current unix timestamp
	// is greater than or equal to the expiration date
	// if so - invalidate the object
	if c.expires > 0 && c.unix >= c.expires {
		c.Valid = false
	}
}
