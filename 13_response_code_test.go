package belajargolang

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)

	res := recorder.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	fmt.Println(string(body))

}
