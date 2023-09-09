package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func Download(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}
	writer.Header().Add("Content-Disposition", "attachment; file=\""+file+"\"")
	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFileServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Download)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
