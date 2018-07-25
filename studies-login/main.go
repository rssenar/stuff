package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/admin/", admin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	url, _ := user.LogoutURL(ctx, "/")
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, `Welcome, %s! <br>`, u.Email)
	fmt.Fprintf(res, `You are admin: %v  <br>`, u.Admin)
	if u.Admin {
		fmt.Fprint(res, `(<a href="/admin">go to admin</a>) <br>`)
	}
	fmt.Fprintf(res, `(<a href="%s">sign out</a>)`, url)
}

func admin(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	url, _ := user.LogoutURL(ctx, "/")
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, `Welcome ADMIN, %s! <br>`, u.Email)
	fmt.Fprintf(res, `(<a href="%s">sign out</a>)`, url)
}
