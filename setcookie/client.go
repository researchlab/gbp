package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	v := url.Values{}
	v.Set("bodykey", "bodyvalue")
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8080/setcookie", body)

	// client  set cookie
	cookie := http.Cookie{Name: "clientcookieid", Value: "121", Expires: time.Now().Add(111 * time.Second)}
	req.AddCookie(&cookie)
	req.AddCookie(&http.Cookie{
		Name:    "clientcookieid2",
		Value:   "id2",
		Expires: time.Now().Add(111 * time.Second),
	})

	//set Content-Type for body post to server normal.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	fmt.Println("req: ", req)

	resp, err := client.Do(req)

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)

	fmt.Println("cookies")
	for _, v := range resp.Cookies() {
		fmt.Printf("%+v\n", v)
	}
}
