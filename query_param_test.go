package belajargolang

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func SayHelloMultiple(writer http.ResponseWriter, request *http.Request) {
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s ", first_name, last_name)
}
func TestQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/hello?name=eko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

	request = httptest.NewRequest(http.MethodGet, "http://localhost/hello?first_name=eko&last_name=khannedy", nil)
	recorder = httptest.NewRecorder()

	SayHelloMultiple(recorder, request)
	res = recorder.Result()
	body, _ = io.ReadAll(res.Body)
	fmt.Println(string(body))

}
