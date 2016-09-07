package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	//data := make(url.Values)
	//data["key"] = []string{"this is key"}
	//data["value"] = []string{"this is value"}

	//把post表单发送给目标服务器
	res, err := http.PostForm("http://127.0.0.1:8091/postpage", url.Values{
		"key":   {"this is url key"},
		"value": {"this is url value"},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()
	fmt.Println("post send success")
}
