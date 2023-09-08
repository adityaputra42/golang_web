package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml"))

	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template data layout",
		"Name":  "Aditya Putra",
		// "Address": map[string]interface{}{
		// 	"Street": "Jalan Sempor Baru",
		// 	"City":   "Kebumen",
		// },
	})
}

func TestTempleateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
