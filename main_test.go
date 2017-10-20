package main

import (
	"os"
	"testing"
)

func TestDefinePort(t *testing.T) {
	port := DefinePort()

	if port != ":8080" {
		t.Fatalf("ENV was not defined but port changed to %s", port)
	}

	os.Setenv("PORT", "8081")
	envPort := DefinePort()

	if envPort != ":8081" {
		t.Fatalf("ENV was defined but port changed to %s", port)
	}
}
