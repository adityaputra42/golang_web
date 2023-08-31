package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}
func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080/hello/", nil)
	request.Header.Add("Content-Type", "aplication/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("x-powered-by", "X-Aditya Putra Pratama-x")

	fmt.Fprint(writer, "OK==>")
}
func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	powered := recorder.Header().Get("x-powered-by")
	fmt.Println(powered)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}
