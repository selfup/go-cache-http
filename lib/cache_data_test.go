package lib

import (
	"reflect"
	"testing"
)

func TestCacheData_NewCacheData(t *testing.T) {
	testStr := "test"
	var (
		unixInt    int64 = 1507351311
		expiresInt int64
	)

	cd := NewCacheData(testStr, unixInt, expiresInt)

	if !reflect.DeepEqual(testStr, cd.data) {
		t.Fatalf("unexpected: %v - actual is: %v", testStr, cd.data)
	}

	if !reflect.DeepEqual(unixInt, cd.unix) {
		t.Fatalf("unexpected: %v - actual is: %v", testStr, cd.data)
	}
}

func TestCacheData_ExpiredCacheData(t *testing.T) {
	testStr := "test"
	var (
		unixInt    int64 = 1507351311
		expiresInt int64 = 1507351310
	)

	cd := NewCacheData(testStr, unixInt, expiresInt)
	cd.ExpirationValidator()

	if true == cd.Valid {
		t.Fatalf("unexpected: %v - actual is: %v", true, cd.Valid)
	}
}
