package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	arg := os.Args[1]
	resp, err := http.Get(arg)
	if err != nil {
		panic(err)
	}
	fmt.Println("--------------------------------")
	fmt.Println("URL TO CHECK :", arg)
	fmt.Println("--------------------------------")
	for k, v := range resp.Header {
		fmt.Printf("resp.Header : %v ==== %v \n", k, v)
	}
	fmt.Println("resp.Trailer                   :", resp.Trailer)
	fmt.Println("resp.Cookies                   :", resp.Cookies())
	fmt.Println("resp.Status                    :", resp.Status)
	fmt.Println("resp.StatusCode                :", resp.StatusCode)
	fmt.Println("resp.Request.Referer()         :", resp.Request.Referer())
	fmt.Println("resp.Request.URL.EscapedPath() :", resp.Request.URL.EscapedPath())
	fmt.Println("resp.Request.URL.ForceQuery    :", resp.Request.URL.ForceQuery)
	fmt.Println("resp.Request.URL.Fragment      :", resp.Request.URL.Fragment)
	fmt.Println("resp.Request.URL.Host          :", resp.Request.URL.Host)
	fmt.Println("resp.Request.URL.IsAbs         :", resp.Request.URL.IsAbs())
	fmt.Println("resp.Request.URL.Opaque        :", resp.Request.URL.Opaque)
	// fmt.Println("resp.Request.URL.Parse         :", resp.Request.URL.Parse)
	fmt.Println("resp.Request.URL.Path          :", resp.Request.URL.Path)
	fmt.Println("resp.Request.URL.Query()       :", resp.Request.URL.Query())
	fmt.Println("resp.Request.URL.RawPath       :", resp.Request.URL.RawPath)
	fmt.Println("resp.Request.URL.RawQuery      :", resp.Request.URL.RawQuery)
	fmt.Println("resp.Request.URL.RequestURI()  :", resp.Request.URL.RequestURI())
	fmt.Println("resp.Request.URL.Scheme        :", resp.Request.URL.Scheme)
	fmt.Println("resp.Request.URL.String()      :", resp.Request.URL.String())
	fmt.Println("resp.Request.URL.User          :", resp.Request.URL.User)
	fmt.Println("resp.Request.UserAgent()       :", resp.Request.UserAgent())
	fmt.Println("resp.Request.TransferEncoding  :", resp.Request.TransferEncoding)
	fmt.Println("resp.Request.Trailer           :", resp.Request.Trailer)
	fmt.Println("resp.Request.RequestURI        :", resp.Request.RequestURI)
	fmt.Println("resp.Request.RemoteAddr        :", resp.Request.RemoteAddr)
	fmt.Println("resp.Request.ProtoMinor        :", resp.Request.ProtoMinor)
	fmt.Println("resp.Request.ProtoMajor        :", resp.Request.ProtoMajor)
	fmt.Println("resp.Request.Proto             :", resp.Request.Proto)
	fmt.Println("resp.Request.PostForm          :", resp.Request.PostForm)
	fmt.Println("resp.Request.ParseForm()       :", resp.Request.ParseForm())
	fmt.Println("resp.Request.MultipartForm     :", resp.Request.MultipartForm)
	fmt.Println("resp.Request.Method            :", resp.Request.Method)
	fmt.Println("resp.Request.Host              :", resp.Request.Host)
	fmt.Println("resp.Request.Header            :", resp.Request.Header)
	fmt.Println("resp.Request.Form              :", resp.Request.Form)
	fmt.Println("resp.Request.Cookies()         :", resp.Request.Cookies())
	fmt.Println("resp.Request.Context()         :", resp.Request.Context())
	fmt.Println("resp.Request.ContentLength     :", resp.Request.ContentLength)

}
