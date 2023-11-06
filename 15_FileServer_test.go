package belajargolang

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileserver := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileSeverGolangEmbed(t *testing.T) {
	// directory := http.Dir("./resources")
	directory, _ := fs.Sub(resources, "resources") // untuk masuk kedalam sub directory
	fileserver := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
