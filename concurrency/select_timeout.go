/**
*     File: select_timeout.go
*	 Brief: 并发超时处理
*	 Descr:
*
*   Author: Hong Li, leehongitrd@163.com
*   Github: https://github.com/researchlab
*  Created: 2016-02-18 16时10分24秒
* Modified: 2016-02-18 16时10分24秒
**/

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://www.grdtechs.com/"
	content, ok := getHttpRes(url)
	fmt.Println("content: ", content, " status:", ok)
}

// 获取url的访问值，返回值:1.成功,返回Body部分,2.失败 返回err 3.超时
func getHttpRes(url string) (string, error) {
	res := make(chan *http.Response, 1)
	httpError := make(chan *error)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			httpError <- &err
		}
		res <- resp
	}()

	for {
		select {
		case r := <-res:
			result, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			return string(result), err
		case err := <-httpError:
			return "err", *err
		case <-time.After(2000 * time.Millisecond):
			return "Timed out", errors.New("Timed out")
		}
	}

}
