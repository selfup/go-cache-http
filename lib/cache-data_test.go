package lib

import (
	"reflect"
	"testing"
)

func TestCacheData_NewCacheData(t *testing.T) {
	testStr := "test"
	var unixInt int64 = 1507351311

	cd := NewCacheData(testStr, unixInt)

	if !reflect.DeepEqual(testStr, cd.data) {
		t.Fatalf("unexpected: %v - actual is: %v", testStr, cd.data)
	}

	if !reflect.DeepEqual(unixInt, cd.unix) {
		t.Fatalf("unexpected: %v - actual is: %v", testStr, cd.data)
	}
}
