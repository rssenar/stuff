package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// upload logic
func upload(path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method:", r.Method)

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		fmt.Println(path + handler.Filename)
		defer f.Close()
		count, err := io.Copy(f, file)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%d bytes uploaded, Name: %s", count, handler.Filename)
		log.Printf("%d bytes uploaded", count)
	})
}

func main() {
	r := mux.NewRouter()
	r.Handle("/upload", upload("/tmp/uploads/")).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", r))
}
