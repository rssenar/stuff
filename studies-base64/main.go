package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "83 71 86 115 98 71 56 115 67 110 100 118 99 109 120 107 73 61 61"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("I'm giving her all she's got Captain!", err)
	}
	fmt.Println(string(bs))
}
