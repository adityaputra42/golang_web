package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CheckHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(400)
		fmt.Fprint(writer, "Name is empty!!")
	} else {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Hello %s", name)
	}
}
func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080", nil)
	recorder := httptest.NewRecorder()

	CheckHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(response.StatusCode)

	fmt.Println(response.Status)

	fmt.Println(bodyString)
}

func TestResponseCodeSucces(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080/?name=Aditya", nil)

	recorder := httptest.NewRecorder()

	CheckHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(response.StatusCode)

	fmt.Println(response.Status)

	fmt.Println(bodyString)
}
