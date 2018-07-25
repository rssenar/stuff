package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// API to shorten url
type apiResponse struct {
	ID, Kind, LongURL string
}

func chkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	longURL := os.Args[len(os.Args)-1]
	body := bytes.NewReader([]byte(fmt.Sprintf(`{"longUrl":"%s"}`, longURL)))
	request, err := http.NewRequest(
		"POST",
		"https://www.googleapis.com/urlshortener/v1/url",
		body)
	chkErr(err)

	request.Header.Add("Content-Type", "application/json")
	q := request.URL.Query()
	q.Add("key", "AIzaSyBdf0GdcEB9Zc80iv49rcEKOMJYWMj_55o")
	request.URL.RawQuery = q.Encode()

	// fmt.Println(request.URL.String())
	client := http.Client{}
	response, err := client.Do(request)
	chkErr(err)

	outputAsBytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var output apiResponse
	err = json.Unmarshal(outputAsBytes, &output)
	chkErr(err)

	// fmt.Println(string(outputAsBytes))
	fmt.Printf("shortened URL: %s\n", output.ID)
	fmt.Println("")
}
