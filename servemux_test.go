package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSeveMux(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/main", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Welcome to Main page")
	})
	mux.HandleFunc("/images/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Images")
	})

	mux.HandleFunc("/images/thumbnail/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
