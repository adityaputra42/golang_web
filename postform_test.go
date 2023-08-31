package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	name := request.PostForm.Get("name")

	age := request.PostForm.Get("age")
	fmt.Fprintf(writer, "nama := %s umur := %s", name, age)

}
func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("name=Aditya&age=27")
	request := httptest.NewRequest(http.MethodPost, "http://locahost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
