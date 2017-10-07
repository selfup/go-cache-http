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

	WriteToState(key, testStr, unixInt, state)

	expected := testStr
	actual := state[key].data

	if !reflect.DeepEqual(testStr, state[key].data) {
		t.Fatalf("unexpected: %v - actual is: %v", expected, actual)
	}
}
