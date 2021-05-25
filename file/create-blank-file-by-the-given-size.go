package main

import (
	"log"
	"os"
)

//var size = int64(1024 * 1e10) // 9.3T
var size = int64(8589934592) // 8 G

// 创建指定大小的空文件
func main() {
	ExTruncate()
	ExSeek()
}

func ExTruncate() {
	f, err := os.Create("foobar1.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := f.Truncate(size); err != nil {
		log.Fatal(err)
	}
}

func ExSeek() {
	f, err := os.Create("foobar2.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Seek(size-1, 0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte{0})
	if err != nil {
		log.Fatal(err)
	}
}
