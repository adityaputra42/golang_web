package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	// bisa diterapkan logic contoh seperti pengecekan login session
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (handler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	handler.Handler.ServeHTTP(writer, request)

}

func TestMiddlewareServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprint(writer, "Hello Middleware")

	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo executed")
		fmt.Fprint(writer, "Hello Foo")

	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic executed")
		panic("ups salah")

	})
	LogMiddleware := new(LogMiddleware)
	LogMiddleware.Handler = mux

	errorHandler := ErrorHandler{Handler: LogMiddleware}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &errorHandler,
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
