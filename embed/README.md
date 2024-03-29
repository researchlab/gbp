Go程序可以使用embed包嵌入静态文件, 打包时也会把嵌入的目录一起打包到二进制文件中, 执行时就不需要在提供嵌入的目录了

```golang
package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
)

// content持有服务器static目录下的所有文件
//go:embed static/*
var content embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(content)))
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	// 读取服务器static目录下的内容
	entries, err := content.ReadDir("static")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		resp, err := http.Get("http://localhost:8080/static/" + e.Name())
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q: %s", e.Name(), body)
	}
}
```
