package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"upload"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is main page")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		fileName := r.Form["name"]
		http.ServeFile(w, r, "file/"+strings.Join(fileName, ""))
	}
}

func main() {
	http.HandleFunc("/upload", http.HandlerFunc(upload.UploadHandler))
	http.HandleFunc("/file", http.HandlerFunc(fileHandler))
	server := &http.Server{
		Addr: ":8080",
	}
	fmt.Println("Server is listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
