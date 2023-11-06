package belajargolang

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
func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Programmer Zaman Now")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-type", "application/json")

	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}

// func TestResponseHeader(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
// 	request.Header.Add("Content-type", "application/json")

// 	recorder := httptest.NewRecorder()
// 	ResponseHeader(recorder, request)

// 	res := recorder.Result()
// 	body, _ := io.ReadAll(res.Body)

// 	fmt.Println(string(body))
// 	fmt.Println(res.Header.Get("X-Powered-by"))

// }
