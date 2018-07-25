package main

import (
	"fmt"

	"github.com/gorilla/securecookie"
)

func main() {
	fmt.Println(fmt.Sprintf("%x", securecookie.GenerateRandomKey(4)))
}
