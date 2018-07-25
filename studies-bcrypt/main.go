package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	password := "GabieisaTuks"
	h := sha256.New()
	h.Write([]byte(password))
	sha := fmt.Sprintf("%x\n", h.Sum(nil))
	fmt.Println(sha[:8])
	// fmt.Printf("%x\n", h.Sum(nil))
}
