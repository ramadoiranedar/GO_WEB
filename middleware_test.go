package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Oops something went wrong... Error detected!")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BEFORE EXECUTE HANDLER")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("AFTER EXECUTE HANDLER")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(w, "Hello foo")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic Executed")
		panic("Oops..errors.")
	})

	// 1. ERROR HANDLER
	// 2. MIDDLEWARE
	// 3. HANDLER MAIN

	logMiddleware := &LogMiddleware{
		Handler: mux, // 3
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware, // 2
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler, // 1
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
