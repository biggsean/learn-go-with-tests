package context1

import (
	"fmt"
	"net/http"
)

// Server is a server
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

// Store interface
type Store interface {
	Fetch() string
	Cancel()
}
