package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	client_cookie, _ := r.Cookie("clientcookieid")

	fmt.Printf("%+v\n", client_cookie)
	for _, v := range r.Cookies() {
		fmt.Printf("%+v\n", v)
	}
	bkey := r.PostFormValue("bodykey")
	fmt.Println(bkey)
	http.SetCookie(w, &http.Cookie{
		Name:    "servercookie",
		Value:   "servercookievalue",
		Expires: time.Now().Add(111 * time.Second),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "servercookie2",
		Value:   "servercookievalue2",
		Expires: time.Now().Add(111 * time.Second),
	})
	io.WriteString(w, "say hi from server")

}
func main() {
	http.HandleFunc("/setcookie", SetCookieHandler)
	fmt.Println("server start at 8080")
	http.ListenAndServe(":8080", nil)
}
