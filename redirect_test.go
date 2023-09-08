package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
}
func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// add logic sebelum redirect
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// add logic sebelum redirect
	http.Redirect(writer, request, "https://www.nusatech.id", http.StatusTemporaryRedirect)
}

func TestRedirectServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
