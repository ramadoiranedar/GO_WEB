package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// HANDLER REQUEST for SERVER
func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// web logic
		_, err := fmt.Fprint(w, "HELLO WORLD!")
		if err != nil {
			panic(err)
		}
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
