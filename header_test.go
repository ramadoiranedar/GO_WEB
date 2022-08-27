package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// REQUEST HEADER
func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	authorization := r.Header.Get("authorization")
	fmt.Fprintln(w, contentType)
	fmt.Fprintln(w, authorization)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhosT:8080", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", "Bearer 0123456789abcdefghijklmnopqrstuvwxyz")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-auth-token", "qwertyuiop")
	fmt.Fprintln(w, "OK")
}

// RESPONSE HEADER
func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhosT:8080", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", "Bearer 0123456789abcdefghijklmnopqrstuvwxyz")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	fmt.Println(response.Header.Get("x-auth-token"))
}
