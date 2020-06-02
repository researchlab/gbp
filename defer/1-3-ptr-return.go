package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func main() {
	data := []byte("golang")
	b := flush(data)
	r, _ := gzip.NewReader(b)
	defer r.Close()
	rawData, err := ioutil.ReadAll(r)
	println(string(rawData))
	if err != nil {
		panic(err.Error())
	}
}

func flush(data []byte) *bytes.Buffer {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(data)
	w.Flush()
	return &b
}
