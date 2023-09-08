package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionMap(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(writer, "if.gohtml", map[string]interface{}{
		"Title": "Template data map",
		// "Name":  "Aditya",
		// "Address": map[string]interface{}{
		// 	"Street": "Jalan Sempor Baru",
		// },
	})
}

func TestTempleateActionMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFiles("./templates/operator.gohtml"))

	t.ExecuteTemplate(writer, "operator.gohtml", map[string]interface{}{
		"Title":      "Template data map",
		"FinalValue": 70,
	})
}

func TestTempleateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template data Range",
		"Hobbies": []string{
			"Game", "Tidur", "Masak", "Makan", "Ngoding",
		},
	})
}

func TestTempleateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "Template data Range",
		"Name":  "Aditya Putra",
		// "Address": map[string]interface{}{
		// 	"Street": "Jalan Sempor Baru",
		// 	"City":   "Kebumen",
		// },
	})
}

func TestTempleateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
