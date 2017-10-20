package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/selfup/gocrashttp/lib"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	FetchCacheOrUpdate(make(map[string]*lib.CacheData))(w, r)
}

func TestItFailsToParseAnEmptyRequest(t *testing.T) {
	server := httptest.NewServer(&myHandler{})
	defer server.Close()

	data := ``
	resp, err := http.NewRequest(
		"POST",
		server.URL,
		strings.NewReader(data),
	)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(resp)

	if res.StatusCode != 500 {
		t.Fatalf("Received non-500 response: %d\n", res.StatusCode)
	}
}

func TestItParsesAValidRequest(t *testing.T) {
	server := httptest.NewServer(&myHandler{})
	defer server.Close()

	data := `{"data": "1", "unix": 12345678}`
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

func TestItRespondsWithA405WhenNotAPostRequest(t *testing.T) {
	server := httptest.NewServer(&myHandler{})
	defer server.Close()

	data := `{}`
	request, err := http.NewRequest(
		"GET",
		server.URL,
		strings.NewReader(data),
	)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 405 {
		t.Fatalf("Received non-405 response: %d\n", res.StatusCode)
	}
}
