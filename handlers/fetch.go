package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/selfup/gocrashttp/lib"
)

// FetchCacheOrUpdate is the main handler
func FetchCacheOrUpdate(state map[string]*lib.CacheData) func(
	w http.ResponseWriter,
	r *http.Request,
) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "this endpoint only supports POST requests", 405)
			return
		}

		decoder := json.NewDecoder(r.Body)

		var incoming struct {
			Key     string
			Data    string
			Unix    int64
			Expires int64
		}

		err := decoder.Decode(&incoming)
		if err != nil {
			http.Error(w, "failed to parse JSON", 500)
			return
		}

		lib.WriteToState(
			incoming.Key,
			incoming.Data,
			incoming.Unix,
			incoming.Expires,
			state,
		)

		outgoing, err := json.Marshal(state[incoming.Key])
		if err != nil {
			http.Error(w, "failed to stringify JSON", 500)
			return
		}

		w.Write(outgoing)
	}
}
