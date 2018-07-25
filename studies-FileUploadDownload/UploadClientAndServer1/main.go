// curl -X POST -H "Content-Type: application/octet-stream" --data-binary '@filename' http://127.0.0.1:5050/upload

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	FileContentType := http.DetectContentType(data)
	w.Write([]byte(FileContentType))

	// file, _ := os.Create("./result.pdf")
	// n, _ := io.Copy(file, r.Body)
	// w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	// go func() {
	// 	time.Sleep(time.Second * 1)
	// 	upload()
	// }()
	fmt.Println("Server running on port:5050")
	http.ListenAndServe(":5050", nil)
}

func upload() {
	file, _ := os.Open("/Users/richardsenar/Documents/Books/beautiful-packages-in-go-london.pdf")
	defer file.Close()
	res, _ := http.Post("http://127.0.0.1:5050/upload", "binary/octet-stream", file)
	message, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println(string(message))
}
