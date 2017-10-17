package lib

import (
	"reflect"
	"testing"
)

func TestCache_WriteToState(t *testing.T) {
	var (
		state         = make(map[string]*CacheData)
		unixInt int64 = 1507351311
	)

	key := "1"
	testStr := "test"

	WriteToState(key, testStr, unixInt, 0, state)

	expected := NewCacheData(testStr, unixInt, 0)
	actual := state[key]

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unexpected: %v - actual is: %v", expected, actual)
	}
}

func TestCache_WriteToStateAndRemoveInvalidKeyValue(t *testing.T) {
	var (
		state            = make(map[string]*CacheData)
		unixInt    int64 = 1507351311
		expiresInt int64 = 1507351310
	)

	key := "1"
	testStr := "test"

	WriteToState(key, testStr, unixInt, expiresInt, state)

	if !(nil == state[key]) {
		t.Fatalf("unexpected: %v - actual is: %v", nil, state[key])
	}
}
