package context1

import (
	"context"
	"fmt"
	"net/http"
)

// Server is a server
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // TODO: log error however you like
		}
		fmt.Fprint(w, data)
	}
}

// Store interface
type Store interface {
	Fetch(ctx context.Context) (string, error)
}
