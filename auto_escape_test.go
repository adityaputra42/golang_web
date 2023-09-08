package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplate.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Iki lho body</p>",
	})
}

func TestAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	AutoEscape(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(AutoEscape),
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func AutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplate.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<h2>Iki lho body</h2>"),
	})
}

func TestAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(AutoEscapeDisabled),
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplate.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>xss</p>", nil)
	recorder := httptest.NewRecorder()
	TemplateXSS(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
