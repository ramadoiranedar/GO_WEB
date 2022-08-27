package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// FORM POST
func FormPost(w http.ResponseWriter, r *http.Request) {
	// manual parsing
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// not manual parsing using built in PostForValue()
	// firstName := r.PostFormValue("first_name")
	// lastName := r.PostFormValue("last_name")

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")
	fmt.Fprintf(w, "%s %s", firstName, lastName)
}

func TestFormTest(t *testing.T) {
	requestBody := strings.NewReader("first_name=Ario&last_name=Damar")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
