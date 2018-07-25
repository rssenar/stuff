package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	file, err := os.Open("/Users/richardsenar/Desktop/AuffenbergFordC403(Mar23-26)15k[12x15MarchMayhem].html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}

	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	// contentType := magic.MIMEType(buffer)
	fmt.Println(contentType)
}
