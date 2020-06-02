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
	"net/http"
	"time"

	"github.com/researchlab/gdk/net"
)

var url = "http://www.baidu.com"

//  main Http Request with select timeout
func main() {
	res, err := HttpGet(&GetOptions{Url: url, Timeout: 10})
	if err != nil {
		return
	}
	siteData := new(interface{})
	err = res.Unmarshal(siteData)
	if err != nil {
		return
	}
}

// GetOptions
type GetOptions struct {
	Url     string `json:"url"`
	Timeout int    `json:"timeout"` // unit(s)
}

// Get Request timeout
var ERR_REQUEST_TIMEOUT = errors.New("Http Request Timeout")

//  HttpGet  http get request with 2s default timeout
func HttpGet(g *GetOptions) (w *net.Response, err error) {
	if g.Timeout <= 0 {
		g.Timeout = 2
	}

	errChan := make(chan error)

	go func() {
		defer func() {
			errChan <- err
		}()
		r, err := http.Get(g.Url)
		if err != nil {
			return
		}
		defer r.Body.Close()
		w, err = net.ResponseRecorder(r)
	}()

	select {
	case err = <-errChan:
	case <-time.After(time.Duration(g.Timeout) * time.Second):
		return nil, ERR_REQUEST_TIMEOUT
	}
	return w, err
}
