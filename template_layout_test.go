package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template Action Layout",
		"Name":  "Damar",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Print(string(body))
}
