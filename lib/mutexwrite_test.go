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

	lidOneIds := state["1"]
	expectedSlice := []string{}
	expected := append(expectedSlice, "1")

	if !reflect.DeepEqual(lidOneIds, expected) {
		t.Fatalf("unexpected slice: %v", expected)
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
