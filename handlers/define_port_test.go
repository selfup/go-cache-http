package handlers

import (
	"os"
	"testing"
)

func TestDefinePort(t *testing.T) {
	port := DefinePort()

	if port != ":8081" {
		t.Fatalf("ENV was not defined but port changed to %s", port)
	}

	os.Setenv("PORT", "8082")
	envPort := DefinePort()

	if envPort != ":8082" {
		t.Fatalf("ENV was defined but port changed to %s", port)
	}
}
