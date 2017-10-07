package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	FetchCacheOrUpdate(w, r)
}

func TestItFailsToParseAnEmptyRequest(t *testing.T) {
	server := httptest.NewServer(&myHandler{})
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 500 {
		t.Fatalf("Received non-500 response: %d\n", resp.StatusCode)
	}
}

func TestItParsesAValidRequest(t *testing.T) {
	server := httptest.NewServer(&myHandler{})
	defer server.Close()
	data := `{"username": "dennis", "balance": 200}`

	request, err := http.NewRequest(
		"POST",
		server.URL,
		strings.NewReader(data),
	)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", res.StatusCode)
	}
}

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
