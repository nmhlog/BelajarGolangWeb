package belajargolang

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_Server(t *testing.T) {

	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello world")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(Writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(Writer, "Hello World")
	})
	mux.HandleFunc("/hi", func(Writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(Writer, "Hi")
	})
	mux.HandleFunc("/images", func(Writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(Writer, "Hi")
	})
	mux.HandleFunc("/images/", func(Writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(Writer, "Hi")
	})
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
