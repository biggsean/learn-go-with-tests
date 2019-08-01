package context1

import (
	"context"
	"net/http"
)

// Server is a server
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pass
	}
}

// Store interface
type Store interface {
	Fetch(ctx context.Context) (string, error)
}
