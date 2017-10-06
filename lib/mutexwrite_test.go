package lib

import (
	"reflect"
	"testing"
)

func TestCache_WriteToState(t *testing.T) {
	var (
		state = make(map[string][]string)
	)

	WriteToState("1", "1", state)
	WriteToState("2", "1", state)

	actual := state["1"]
	expected := []string{"1", "2"}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("unexpected: %v - actual is: %v", expected, actual)
	}
}

func TestCache_WriteToLidIds_NoWrites(t *testing.T) {
	var (
		state = make(map[string][]string)
	)

	mockSlice := []string{"1"}

	WriteToLidIds(mockSlice, "1", "1", state)

	if reflect.DeepEqual(mockSlice, state["1"]) {
		t.Fatalf("unexpected slice: %v", state)
	}
}

func TestCache_WriteToLidIds_Writes(t *testing.T) {
	var (
		state = make(map[string][]string)
	)

	mockSlice := []string{"1"}

	WriteToLidIds(mockSlice, "1", "2", state)

	if reflect.DeepEqual(mockSlice, state["2"]) {
		t.Fatalf("unexpected slice: %v", state)
	}
}
