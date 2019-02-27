package upload

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getFileID() int {
	file, err := os.Open("upload/fileCounter.txt")
	checkErr(err)
	defer file.Close()

	var counter int
	fmt.Fscanf(file, "%d", &counter)
	return counter + 1
}

func updateFileCounter(n int) {
	file, err := os.OpenFile("upload/fileCounter.txt", os.O_WRONLY, 0222)
	checkErr(err)
	defer file.Close()

	fmt.Fprintf(file, "%d", n)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("upload/upload.gtpl")
	checkErr(err)

	if r.Method == "GET" {
		t.Execute(w, "")
	}

	if r.Method == "POST" {
		r.ParseMultipartForm(1 << 32)
		file, fileHeader, err := r.FormFile("userFile")
		checkErr(err)
		defer file.Close()

		newFileID := getFileID()
		newFileName := strconv.Itoa(newFileID) + "_" + template.HTMLEscapeString(fileHeader.Filename)

		newFile, err := os.OpenFile("file/"+newFileName, os.O_WRONLY|os.O_CREATE, 0666)

		checkErr(err)
		defer newFile.Close()

		io.Copy(newFile, file)

		updateFileCounter(newFileID)

		t.Execute(w, "http://localhost:8080/file?name="+newFileName)
	}
}
