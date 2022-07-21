package main

import (
	"encoding/json"
	"net/http"
)

type File struct {
	Name string
	Type string
	Size int
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./views/")))
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8090", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	_, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "error: not such file", http.StatusBadRequest)
		return
	}

	fileBytes, err := json.Marshal(File{
		Name: handler.Filename,
		Type: handler.Header["Content-Type"][0],
		Size: int(handler.Size),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fileBytes)
}
