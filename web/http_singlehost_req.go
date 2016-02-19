/**
*     File: http_singlehost_req.go
*	 Brief: 通过自定义路由规则只响应指定host请求
*	 Descr:
*
*   Author: Hong Li, leehongitrd@163.com
*   Github: https://github.com/researchlab
*  Created: 2016-02-19 09时13分45秒
* Modified: 2016-02-19 09时13分45秒
**/

package main

import (

	//"fmt"
	//"io"
	"net/http"
)

//限制域名请求
func SingleHost(handler http.Handler, allowhost string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Host == allowhost {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(fn)
}

func light(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "好吧我的肚子好像饿了，这边是利用函数实现了http中间件，而不是用一个类型，函数可以是一个类型，在函数当中做判断")
	w.Write([]byte("肚子饿了"))
}

func main() {
	single := SingleHost(http.HandlerFunc(light), "localhost:8080")
	http.ListenAndServe(":8080", single)
}
