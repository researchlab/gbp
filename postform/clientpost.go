package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	client := &http.Client{}
	res, err := client.PostForm("http://127.0.0.1:8091/postpage", url.Values{
		"key":   {"this is client key"},
		"value": {"this is client value"},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()
	fmt.Println("post send success")

}
