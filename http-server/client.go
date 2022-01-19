package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	update := flag.Bool("update", false, "update domain info")
	stock := flag.Bool("stock", false, "stock info")
	flag.Parse()
	fmt.Println("update:", *update)
	if *update {
		httpPost()
	} else if *stock {
		httpStock()
	} else {
		httpGet()
	}
}

func httpPost() {
	di := &DomainInfo{
		Name:  "name-" + time.Now().String(),
		SrcIp: "192.168.1.1",
	}
	b, err := json.Marshal(di)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://127.0.0.1:8000/domain/update", "application/x-www-form-urlencoded", strings.NewReader(fmt.Sprintf("param=%s", b)))
	if err != nil {
		fmt.Println("ResponseError:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func httpGet() {
	resp, err := http.Get("http://127.0.0.1:8000/domain/read")
	if err != nil {
		fmt.Println("ResponseError:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func httpStock() {
	resp, err := http.Get("http://127.0.0.1:8000/stock/api/domain/read?code=10011")
	//resp, err := http.Get("http://127.0.0.1:8000/stock/api/domain/read")
	if err != nil {
		fmt.Println("ResponseError:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
