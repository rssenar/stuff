package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// call upload on route
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		/*
		   // error handling here
		   // return since file upload filled
		*/
		return
	}
	defer file.Close()

	if handler.Header["Content-Type"][0] != "text/html" {
		/*
		   // error handling here
		   // Close file explicitly now might not be needed but I always do
		   // return since file upload filled
		*/
		file.Close()
		return
	}
	// set route to file location to be saved
	f, err := os.OpenFile("./public/assets/emails/"+handler.Filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		/*
		   // error handling here
		   // return since file upload filled
		*/
		return
	}
	defer f.Close()
	// add file infomation: name, path, and file size to database optional
	// models.AddTemplate(handler.Filename, "./public/assets/emails/"+handler.Filename, strconv.Itoa(int(handler.Size)))
	// copy file to file system
	io.Copy(f, file)
	// redirect to whereever
	http.Redirect(w, r, "/dashboard", 302)
	// return after redirect
	return
}
