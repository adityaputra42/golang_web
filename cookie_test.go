package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-APP-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintln(writer, "Sukses set cookie")
}
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-APP-Name")
	if err != nil {
		fmt.Fprintln(writer, "No Cookie")
	} else {
		fmt.Fprintln(writer, cookie.Value)
	}

}
func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set_cookie", SetCookie)
	mux.HandleFunc("/get_cookie", GetCookie)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080/?name=Aditya%20Putra%20Pratama", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Println(cookie.Value)
	}

}
func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-APP-Name"
	cookie.Value = "Gita Prigi"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(response.StatusCode)

	fmt.Println(response.Status)

	fmt.Println(bodyString)

}
